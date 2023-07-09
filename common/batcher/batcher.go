package batcher

// 批量数据聚合
import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"sync"
	"time"
)

// msg 是用于传递 key-value 数据的结构体
type msg struct {
	key string
	val interface{}
}

// Options 是 Batcher 的配置选项
type Options struct {
	Worker   int           // 工作协程数
	Buffer   int           // 缓冲区大小
	Size     int           // 批处理大小
	Interval time.Duration // 批处理间隔（单位 s）
}

// Batcher 是一个批处理器
type Batcher struct {
	opts Options

	Do       func(ctx context.Context, val map[string][]interface{}) // 处理批处理数据的函数
	Sharding func(key string) int                                    // 用于根据 key 进行分片的函数
	chans    []chan *msg                                             // 批处理通道数组
	wait     sync.WaitGroup                                          // 等待所有批处理协程完成的 WaitGroup
}

// New 创建一个 Batcher 实例
func New(o Options) *Batcher {
	b := &Batcher{}
	b.opts = o
	b.chans = make([]chan *msg, b.opts.Worker)
	for i := 0; i < b.opts.Worker; i++ {
		b.chans[i] = make(chan *msg, b.opts.Buffer)
	}
	return b
}

// Start 启动批处理器
func (b *Batcher) Start() {
	if b.Do == nil {
		log.Fatal("Batcher: Do func is nil")
	}
	if b.Sharding == nil {
		log.Fatal("Batcher: Sharding func is nil")
	}
	b.wait.Add(len(b.chans))

	for i, ch := range b.chans {
		go b.merge(i, ch)
	}
}

// Add 添加一条 key-value 数据到批处理器
func (b *Batcher) Add(key string, val interface{}) error {
	ch, msg := b.add(key, val)
	select {
	case ch <- msg:
	default:
		return errors.New("Err Full")
	}
	return nil
}

// add 将 key-value 数据添加到相应的批处理通道中
func (b *Batcher) add(key string, val interface{}) (chan *msg, *msg) {
	sharding := b.Sharding(key) % b.opts.Worker
	ch := b.chans[sharding]
	msg := &msg{key: key, val: val}
	return ch, msg
}

// 两个触发执行Do方法的条件，一是当聚合的数据条数大于等于设置的条数，二是当触发设置的定时器
func (b *Batcher) merge(idx int, ch <-chan *msg) {
	defer b.wait.Done()

	var (
		msg        *msg                                          // 用于存储接收到的消息
		count      int                                           // 计数器，记录已处理的消息数量
		closed     bool                                          // 标记通道是否已关闭
		lastTicker = true                                        // 标记是否是最后一个定时器
		interval   = b.opts.Interval                             // 批处理间隔
		vals       = make(map[string][]interface{}, b.opts.Size) // 存储按 key 分组的值
	)

	// 根据协程的索引设置对应的间隔时间
	// 如果 idx 大于 0，则表示当前协程不是第一个协程，通过调整定时器的间隔时间来平均分配批处理操作。
	// 不同工作协程之间的批处理操作时间错开，可以减少并发操作的冲突和资源竞争
	if idx > 0 {
		interval = time.Duration(int64(idx) * (int64(b.opts.Interval) / int64(b.opts.Worker)))
	}
	fmt.Println(interval)
	ticker := time.NewTicker(interval) // 创建定时器
	for {
		select {
		case msg = <-ch: // 接收到消息
			if msg == nil { // 如果为 nil，表示通道已关闭
				closed = true
				break
			}
			count++
			// 将值添加到按 key 分组的 map 中
			vals[msg.key] = append(vals[msg.key], msg.val)
			// 达到批处理大小
			if count >= b.opts.Size {
				break
			}
			continue
		case <-ticker.C:
			// 如果是最后一个定时器
			if lastTicker {
				ticker.Stop()
				ticker = time.NewTicker(b.opts.Interval)
				// 更新标记
				lastTicker = false
			}
		}

		// 执行批处理操作
		if len(vals) > 0 {
			ctx := context.Background()
			// 调用处理批处理数据的函数
			b.Do(ctx, vals)
			// 重置按 key 分组的 map
			vals = make(map[string][]interface{}, b.opts.Size)
			// 重置计数器
			count = 0
		}

		// 如果通道已关闭，则停止定时器并返回
		if closed {
			ticker.Stop()
			return
		}
	}
}

// Close 关闭批处理器
func (b *Batcher) Close() {
	for _, ch := range b.chans {
		ch <- nil
	}
	b.wait.Wait()
}

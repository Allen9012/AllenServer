package profiler

import (
	"container/list"
	"fmt"
	"github.com/Allen9012/AllenGame/log"
	"sync"
	"time"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/10
  @desc: 性能分析和监控
  @modified by:
**/

// DefaultMaxOvertime 最大超长时间，一般可以认为是死锁或者死循环，或者极差的性能问题
var DefaultMaxOvertime time.Duration = 5 * time.Second

// DefaultOvertime 超过该时间将会监控报告
var DefaultOvertime time.Duration = 10 * time.Millisecond
var DefaultMaxRecordNum int = 100 //最大记录条数
// 分析注册
var mapProfiler map[string]*Profiler

type ReportFunType func(name string, callNum int, costTime time.Duration, record *list.List)

var reportFunc ReportFunType = DefaultReportFunction

type Element struct {
	tagName  string
	pushTime time.Time
}

type RecordType int

const (
	MaxOvertimeType = 1
	OvertimeType    = 2
)

type Record struct {
	RType      RecordType
	CostTime   time.Duration
	RecordName string
}

type Analyzer struct {
	elem     *list.Element
	profiler *Profiler
}

type Profiler struct {
	stack       *list.List //Element
	stackLocker sync.RWMutex
	mapAnalyzer map[*list.Element]Analyzer
	record      *list.List //Record

	callNum       int           //调用次数
	totalCostTime time.Duration //总消费时间长

	maxOverTime  time.Duration
	overTime     time.Duration
	maxRecordNum int
}

func init() {
	mapProfiler = map[string]*Profiler{}
}

func RegProfiler(profilerName string) *Profiler {
	if _, ok := mapProfiler[profilerName]; ok == true {
		return nil
	}

	pProfiler := &Profiler{stack: list.New(), record: list.New(), maxOverTime: DefaultMaxOvertime, overTime: DefaultOvertime}
	mapProfiler[profilerName] = pProfiler
	return pProfiler
}

/* 设置字段 */

func (slf *Profiler) SetMaxOverTime(tm time.Duration) {
	slf.maxOverTime = tm
}

func (slf *Profiler) SetOverTime(tm time.Duration) {
	slf.overTime = tm
}

func (slf *Profiler) SetMaxRecordNum(num int) {
	slf.maxRecordNum = num
}

func (slf *Profiler) Push(tag string) *Analyzer {
	slf.stackLocker.Lock()
	defer slf.stackLocker.Unlock()

	pElem := slf.stack.PushBack(&Element{tagName: tag, pushTime: time.Now()})

	return &Analyzer{elem: pElem, profiler: slf}
}

func (slf *Profiler) check(pElem *Element) (*Record, time.Duration) {
	//TODO implement me
	panic("implement me")
}

func (slf *Analyzer) Pop() {
	//TODO implement me
	panic("implement me")
}

func (slf *Profiler) pushRecordLog(record *Record) {
	//TODO implement me
	panic("implement me")
}

func SetReportFunction(reportFun ReportFunType) {
	reportFunc = reportFun
}

func DefaultReportFunction(name string, callNum int, costTime time.Duration, record *list.List) {
	if record.Len() <= 0 {
		return
	}

	var strReport string
	strReport = "Profiler report tag " + name + ":\n"
	var average int64
	if callNum > 0 {
		average = costTime.Milliseconds() / int64(callNum)
	}

	strReport += fmt.Sprintf("process count %d,take time %d Milliseconds,average %d Milliseconds/per.\n", callNum, costTime.Milliseconds(), average)
	elem := record.Front()
	var strTypes string
	for elem != nil {
		pRecord := elem.Value.(*Record)
		if pRecord.RType == MaxOvertimeType {
			strTypes = "too slow process"
		} else {
			strTypes = "slow process"
		}

		strReport += fmt.Sprintf("%s:%s is take %d Milliseconds\n", strTypes, pRecord.RecordName, pRecord.CostTime.Milliseconds())
		elem = elem.Next()
	}

	log.SInfo("report", strReport)
}

func Report() {
	var record *list.List
	for name, prof := range mapProfiler {
		prof.stackLocker.RLock()

		//取栈顶，是否存在异常MaxOverTime数据
		pElem := prof.stack.Back()
		for pElem != nil {
			pElement := pElem.Value.(*Element)
			pExceptionElem, _ := prof.check(pElement)
			if pExceptionElem != nil {
				prof.pushRecordLog(pExceptionElem)
			}
			pElem = pElem.Prev()
		}

		if prof.record.Len() == 0 {
			prof.stackLocker.RUnlock()
			continue
		}

		record = prof.record
		prof.record = list.New()
		callNum := prof.callNum
		totalCostTime := prof.totalCostTime
		prof.stackLocker.RUnlock()

		DefaultReportFunction(name, callNum, totalCostTime, record)
	}
}

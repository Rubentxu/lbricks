package bgo

import "time"
/**
   * Decorator is the base class for all decorator nodes. Thus, if you want to
   * create new custom decorator nodes, you need to inherit from this class.
* */
type Decorator struct {
	BaseNode
	child Node
}

/**
   * The Inverter decorator inverts the result of the child, returning `SUCCESS`
   * for `FAILURE` and `FAILURE` for `SUCCESS`.
  **/
type Inverter struct {
	Decorator
}

func (this *Inverter) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}
	status := ExecuteNode(this.child,context)

	if (status == SUCCESS) {
		status = FAILURE
	} else if (status == FAILURE) {
		status = SUCCESS
	}
	return status

}

func NewInverter(title string,child Node) *Inverter {
	inverter := &Inverter{}
	inverter.ID = CreateUUID()
	inverter.Category = DECORATOR
	inverter.Name = "Inverter"
	inverter.Title = title
	inverter.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return inverter
}

/**
   * This decorator limit the number of times its child can be called. After a
   * certain number of times, the Limiter decorator returns `FAILURE` without
   * executing the child.
 **/
type Limiter struct {
	Decorator
	maxLoop		int16
}

func (this *Limiter) Open(context *Context) {
	context.Blackboard.Set("count", 0 ,context.Tree.Id, this.ID)
}

func (this *Limiter) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count,_ := context.Blackboard.Get("count", context.Tree.Id, this.ID)

	if(count.(int16) < this.maxLoop) {
		status := ExecuteNode(this.child,context)
		if status == SUCCESS || status == FAILURE {
			context.Blackboard.Set("count", count+1 ,context.Tree.Id, this.ID)
		}
		return status
	}
	return FAILURE

}

func NewLimiter(title string,child Node, maxLoop int16) *Limiter {
	limiter := &Limiter{}
	limiter.ID = CreateUUID()
	limiter.Category = DECORATOR
	limiter.Name = "Limiter"
	limiter.Title = title
	limiter.maxLoop = maxLoop
	limiter.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return limiter
}

/**
   * The MaxTime decorator limits the maximum time the node child can execute.
   * Notice that it does not interrupt the execution itself (i.e., the child
   * must be non-preemptive), it only interrupts the node after a `RUNNING`
   * status.
**/
type MaxTime struct {
	Decorator
	maxTime		time.Duration
}

func (this *MaxTime) Open(context *Context) {
	startTime := time.Now()
	context.Blackboard.Set("starTime", startTime ,context.Tree.Id, this.ID)
}

func (this *MaxTime) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	startTime ,_:= context.Blackboard.Get("starTime", context.Tree.Id, this.ID)
	status := ExecuteNode(this.child,context)

	if(time.Since(startTime.(time.Time)) > this.maxTime) {
		return FAILURE

	}
	return status

}

func NewMaxTime(title string, child Node, duration time.Duration) *MaxTime {
	maxTime := &MaxTime{}
	maxTime.ID = CreateUUID()
	maxTime.Category = DECORATOR
	maxTime.Name = "MaxTime"
	maxTime.Title = title
	maxTime.maxTime = duration
	maxTime.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return maxTime
}


/**
   * Repeater is a decorator that repeats the tick signal until the child node
   * return `RUNNING` or `ERROR`. Optionally, a maximum number of repetitions
   * can be defined.
**/
type Repeater struct {
	Decorator
	maxLoop		int8
}

func (this *Repeater) Open(context *Context) {
	context.Blackboard.Set("count", 0 ,context.Tree.Id, this.ID)
}

func (this *Repeater) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count ,_:= context.Blackboard.Get("count", context.Tree.Id, this.ID)
	status := SUCCESS

	for this.maxLoop < 0 || count < this.maxLoop {
		status = ExecuteNode(this.child,context)
		if status == SUCCESS || status == FAILURE {
			count++
		} else {
			break
		}
	}
	context.Blackboard.Set("count", count, context.Tree.Id, this.ID)
	return status

}

func NewRepeater(title string, child Node, maxLoop int8) *Repeater {
	repeater := &Repeater{}
	repeater.ID = CreateUUID()
	repeater.Category = DECORATOR
	repeater.Name = "Repeater"
	repeater.Title = title
	repeater.maxLoop = maxLoop
	repeater.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return repeater
}
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
	inverter.child = child
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
	maxLoop		int
}

func (this *Limiter) Open(context *Context) {
	context.GetNodeMemory(this).Integer["count"] = 0
}

func (this *Limiter) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count := context.GetNodeMemory(this).Integer["count"]

	if(count < this.maxLoop) {
		status := ExecuteNode(this.child,context)
		if status == SUCCESS || status == FAILURE {
			context.GetNodeMemory(this).Integer["count"] =  count+1
		}
		return status
	}
	return FAILURE

}

func NewLimiter(title string, maxLoop int, child Node,) *Limiter {
	limiter := &Limiter{}
	limiter.ID = CreateUUID()
	limiter.Category = DECORATOR
	limiter.Name = "Limiter"
	limiter.Title = title
	limiter.maxLoop = maxLoop
	limiter.child = child
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
	context.GetNodeMemory(this).Time["starTime"]= startTime
}

func (this *MaxTime) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	startTime ,_:= context.GetNodeMemory(this).Time["starTime"]
	status := ExecuteNode(this.child,context)

	if(time.Since(startTime) > this.maxTime) {
		return FAILURE

	}
	return status

}

func NewMaxTime(title string, duration time.Duration, child Node) *MaxTime {
	maxTime := &MaxTime{}
	maxTime.ID = CreateUUID()
	maxTime.Category = DECORATOR
	maxTime.Name = "MaxTime"
	maxTime.Title = title
	maxTime.maxTime = duration * time.Millisecond
	maxTime.child = child
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
	maxLoop		int
}

func (this *Repeater) Open(context *Context) {
	context.GetNodeMemory(this).Integer["count"] = 0
}

func (this *Repeater) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count ,_:= context.GetNodeMemory(this).Integer["count"]
	status := SUCCESS

	for this.maxLoop < 0 || count < this.maxLoop {
		status = ExecuteNode(this.child,context)
		if status == SUCCESS || status == FAILURE {
			count =count +1
		} else {
			break
		}
	}
	context.GetNodeMemory(this).Integer["count"] = count
	return status

}

func NewRepeater(title string, maxLoop int, child Node) *Repeater {
	repeater := &Repeater{}
	repeater.ID = CreateUUID()
	repeater.Category = DECORATOR
	repeater.Name = "Repeater"
	repeater.Title = title
	repeater.maxLoop = maxLoop
	repeater.child = child
	repeater.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return repeater
}


type RepeatUntilFailure struct {
	Decorator
	maxLoop		int
}

func (this *RepeatUntilFailure) Open(context *Context) {
	context.GetNodeMemory(this).Integer["count"] = 0
}

func (this *RepeatUntilFailure) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count ,_:= context.GetNodeMemory(this).Integer["count"]
	status := ERROR

	for this.maxLoop < 0 || count < this.maxLoop {
		status = ExecuteNode(this.child,context)
		if status == SUCCESS {
			count =count +1
		} else {
			break
		}
	}
	context.GetNodeMemory(this).Integer["count"] = count
	return status

}

func NewRepeatUntilFailure(title string, maxLoop int, child Node) *RepeatUntilFailure {
	repeater := &RepeatUntilFailure{}
	repeater.ID = CreateUUID()
	repeater.Category = DECORATOR
	repeater.Name = "RepeatUntilFailure"
	repeater.Title = title
	repeater.maxLoop = maxLoop
	repeater.child = child
	repeater.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return repeater
}



type RepeatUntilSuccess struct {
	Decorator
	maxLoop		int
}

func (this *RepeatUntilSuccess) Open(context *Context) {
	context.GetNodeMemory(this).Integer["count"] = 0
}

func (this *RepeatUntilSuccess) Tick(context *Context) Status {
	if this.child == nil {
		return ERROR
	}

	count ,_:= context.GetNodeMemory(this).Integer["count"]
	status := ERROR

	for this.maxLoop < 0 || count < this.maxLoop {
		status = ExecuteNode(this.child,context)
		if status == FAILURE {
			count =count +1
		} else {
			break
		}
	}
	context.GetNodeMemory(this).Integer["count"] = count
	return status

}

func NewRepeatUntilSuccess(title string, maxLoop int, child Node) *RepeatUntilSuccess {
	repeater := &RepeatUntilSuccess{}
	repeater.ID = CreateUUID()
	repeater.Category = DECORATOR
	repeater.Name = "RepeatUntilSuccess"
	repeater.Title = title
	repeater.maxLoop = maxLoop
	repeater.child = child
	repeater.Description = "Decorator is the base class for all decorator nodes. Thus, if you want to create new custom decorator nodes, you need to inherit from this class. "
	return repeater
}
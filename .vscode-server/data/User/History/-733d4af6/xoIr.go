package worker

type Context interface{
	func function_name();
	func execute(*Context);	
}
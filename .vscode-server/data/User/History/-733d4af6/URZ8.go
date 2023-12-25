package worker



type Context interface{
	func read_block();
	func write_file();
	func task_info();
	func datasets_by_tags();
}

type Worker interface{
	func function_name();
	func execute(*Context);	}
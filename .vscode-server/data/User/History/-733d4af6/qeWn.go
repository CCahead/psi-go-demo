package worker



type Context interface{
	func read_block(); //return Result<Vec<u8>>
	func write_file(); //return Result<()>
	func task_info();   // return &Task
	func datasets_by_tags();// return Vec<DataSet>
}

type Worker interface{
	func function_name();
	func execute(*Context);	}
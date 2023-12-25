package worker


type Result[R res, E error] struct{
	Res R
	Err E
}


type Context interface{
	read_block() []Result //return Result<Vec<u8>>
	write_file(); //return Result<()>
	task_info();   // return &Task
	datasets_by_tags();// return Vec<DataSet>
}

type Worker interface{
	func function_name();

	// return Result<String> This is the thing I need to implement

	func execute();	

	// return Result<String> This is the thing I need to implement

}
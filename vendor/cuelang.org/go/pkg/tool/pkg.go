// Code generated by cuelang.org/go/pkg/gen. DO NOT EDIT.

// Package tool defines stateful operation types for cue commands.
//
// This package is only visible in cue files with a _tool.cue or _tool_test.cue
// ending.
//
// CUE configuration files are not influenced by and do not influence anything
// outside the configuration itself: they are hermetic. Tools solve
// two problems: allow outside values such as environment variables,
// file or web contents, random generators etc. to influence configuration,
// and allow configuration to be actionable from within the tooling itself.
// Separating these concerns makes it clear to user when outside influences are
// in play and the tool definition can be strict about what is allowed.
//
// Tools are defined in files ending with _tool.cue. These files have a
// top-level map, "command", which defines all the tools made available through
// the cue command.
//
// The following definitions are for defining commands in tool files:
//
//	// A Command specifies a user-defined command.
//	//
//	// Descriptions are derived from the doc comment, if they are not provided
//	// structurally, using the following format:
//	//
//	//    // short description on one line
//	//    //
//	//    // Usage: <name> usage (optional)
//	//    //
//	//    // long description covering the remainder of the doc comment.
//	//
//	Command: {
//		// Tasks specifies the things to run to complete a command. Tasks are
//		// typically underspecified and completed by the particular internal
//		// handler that is running them. Tasks can be a single task, or a full
//		// hierarchy of tasks.
//		//
//		// Tasks that depend on the output of other tasks are run after such tasks.
//		// Use `$after` if a task needs to run after another task but does not
//		// otherwise depend on its output.
//		Tasks
//
//		// $usage summarizes how a command takes arguments.
//		//
//		// Example:
//		//     mycmd [-n] names
//		$usage?: string
//
//		// $short is short description of what the command does.
//		$short?: string
//
//		// $long is a longer description that spans multiple lines and
//		// likely contain examples of usage of the command.
//		$long?: string
//	}
//
//	// TODO:
//	// - child commands?
//
//	// Tasks defines a hierarchy of tasks. A command completes if all tasks have
//	// run to completion.
//	Tasks: Task | {
//		[name=Name]: Tasks
//	}
//
//	// #Name defines a valid task or command name.
//	Name: =~#"^\PL([-](\PL|\PN))*$"#
//
//	// A Task defines a step in the execution of a command.
//	Task: {
//		$type: "tool.Task" // legacy field 'kind' still supported for now.
//
//		// $id indicates the operation to run. It must be of the form
//		// packagePath.Operation.
//		$id: =~#"\."#
//
//		// $after can be used to specify a task is run after another one, when
//		// it does not otherwise refer to an output of that task.
//		$after?: Task | [...Task]
//	}
//
//	// TODO: consider these options:
//	//   $success: bool
//	//   $runif: a.b.$success or $guard: a.b.$success
//	// With this `$after: a.b` would just be a shorthand for `$guard: a.b.$success`.
package tool
import React from "react"
import TextField from "material-ui/TextField"
import RaisedButton from "material-ui/RaisedButton"

export default class TaskAddingForm extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			task: "",
		}
	}
	inputTask = e => {
		this.setState({ task: e.target.value })
	}
	submit = () => {
		console.log(this.state.task, " をしようね！")
		this.setState({ task: "" })
		//ここにタスクのPOST
	}
	render() {
		return (
			<div>
				<TextField
					label="Your task"
					value={this.state.task}
					onChange={this.inputTask}
					placeholder="e.g. 進捗を出す"
					margin="normal"
				/>
				<RaisedButton label="タスク追加" onClick={this.submit.bind(this)} />
			</div>
		)
	}
}

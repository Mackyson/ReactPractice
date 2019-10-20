import React from "react"
import { TextField, Button } from "@material-ui/core/"

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
				<TextField value={this.state.task} onChange={this.inputTask} placeholder="e.g. 進捗を出す" margin="none" />
				<Button variant="contained" onClick={this.submit.bind(this)}>
					タスク追加
				</Button>
			</div>
		)
	}
}

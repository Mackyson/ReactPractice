import React from "react"
import { TextField, Button } from "@material-ui/core/"

export default class TaskAddingForm extends React.Component {
	constructor(props) {
		super(props)
		let nowDate = new Date().toISOString().split("T")[0] //yyyy-mm-dd
		this.state = {
			date: nowDate,
			task: "",
			uid: -1,
		}
	}
	inputContent = e => {
		this.setState({ task: e.target.value })
	}
	inputDate = e => {
		this.setState({ date: e.target.value })
	}
	submit = () => {
		console.log(this.state.task, " をしようね！")
		this.setState({ task: "" })
		//ここにタスクのPOST
	}
	render() {
		return (
			<div>
				<TextField value={this.state.task} onChange={this.inputContent} placeholder="e.g. 進捗を出す" margin="none" />
				<TextField
					id="date"
					type="date"
					onChange={this.inputDate}
					value={this.state.date}
					InputLabelProps={{
						shrink: true,
					}}
				/>
				<Button variant="contained" onClick={this.submit.bind(this)}>
					タスク追加
				</Button>
			</div>
		)
	}
}

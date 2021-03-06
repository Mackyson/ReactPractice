import React from "react"
import axios from "axios"
import PropTypes from "prop-types"
import { TextField, Button } from "@material-ui/core/"

export default class TaskAddingForm extends React.Component {
	static propTypes = {
		fetch: PropTypes.func.isRequired,
	}
	constructor(props) {
		super(props)
		let nowDate = new Date().toISOString().split("T")[0] //yyyy-mm-dd
		this.state = {
			date: nowDate,
			content: "",
			uid: 256,
		}
	}
	inputContent = e => {
		this.setState({ content: e.target.value })
	}
	inputDate = e => {
		this.setState({ date: e.target.value })
	}
	submit = () => {
		axios
			.post("http://localhost:8080/todo/" + this.state.uid + "/", {
				content: this.state.content,
				deadline: this.state.date,
			})
			.then(() => {
				this.setState({ content: "" })
				this.props.fetch()
			})
			.catch(error => {
				alert(error)
			})
	}
	render() {
		return (
			<div>
				<TextField
					value={this.state.content}
					onChange={this.inputContent}
					placeholder="e.g. 進捗を出す"
					margin="none"
				/>
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

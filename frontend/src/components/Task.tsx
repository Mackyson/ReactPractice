import React from "react"
import axios from "axios"
import Table from "@material-ui/core/Table"
import TableBody from "@material-ui/core/TableBody"
import TableCell from "@material-ui/core/TableCell"
import TableRow from "@material-ui/core/TableRow"

import TaskAddingForm from "./TaskAddingForm.tsx"
import TaskList from "./TaskList.tsx"

export default class Task extends React.Component {
	constructor(props) {
		super(props)
		this.state = { table: <div>読み込み中</div> }
	}
	fetch = () => {
		axios
			.get("http://localhost:8080/todo/256/", {})
			.then(results => {
				this.setState({ tasks: results.data })
				let cnt = 0
				this.setState({
					table: (
						<Table>
							<TableBody>
								{results.data.map(task => (
									<TableRow key={cnt++}>
										<TableCell>{task.content}</TableCell>
									</TableRow>
								))}
							</TableBody>
						</Table>
					),
				})
			})
			.catch(error => {
				console.log(error)
			})
	}
	render() {
		return (
			<div>
				<TaskAddingForm fetch={this.fetch.bind(this)} />
				<TaskList table={this.state.table} fetch={this.fetch.bind(this)} />
			</div>
		)
	}
}

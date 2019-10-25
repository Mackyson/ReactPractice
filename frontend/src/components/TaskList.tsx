import React from "react"
import axios from "axios"
import Table from "@material-ui/core/Table"
import TableBody from "@material-ui/core/TableBody"
import TableCell from "@material-ui/core/TableCell"
import TableRow from "@material-ui/core/TableRow"

export default class TaskList extends React.Component {
	constructor(props) {
		super(props)
		this.state = { tasks: {}, table: <div>読み込み中</div> }
	}
	fetch = () => {
		//ここにタスクのGET
		axios
			.get("http://localhost:8080/todo/256/", {})
			.then(results => {
				console.log(results.data)
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
	componentDidMount() {
		this.fetch.bind(this)()
	}
	render() {
		return this.state.table
	}
}

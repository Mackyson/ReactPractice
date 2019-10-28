import React from "react"
import axios from "axios"
import qs from "qs"
import { TextField, Button } from "@material-ui/core/"

export default class SignUpForm extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			userName: "",
			password: "",
		}
	}
	inputUsername = e => {
		this.setState({ userName: e.target.value })
	}
	inputPassword = e => {
		this.setState({ password: e.target.value })
	}
	submit = () => {
		axios
			.post("http://localhost:8080/signin", qs.stringify({ name: this.state.userName, password: this.state.password }))
			.then(res => {
				console.log(res)
			})
			.catch(err => {
				console.log(err.response.data.error)
				alert(err.response.data.error)
			})
	}
	render() {
		return (
			<div>
				<TextField label="Your unique name" value={this.state.userName} onChange={this.inputUsername} />
				<TextField label="Your password" type="password" value={this.state.password} onChange={this.inputPassword} />
				<Button variant="contained" onClick={this.submit.bind(this)}>
					ログイン
				</Button>
			</div>
		)
	}
}

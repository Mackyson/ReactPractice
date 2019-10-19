import React from "react"
import RaisedButton from "material-ui/RaisedButton"

export default class SignInForm extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			cnt: 0,
		}
	}
	handleClick = () => {
		this.setState({ cnt: this.state.cnt + 1 })
	}
	render() {
		return (
			<div>
				ログイン画面
				<RaisedButton label={this.state.cnt} onClick={this.handleClick} />
			</div>
		)
	}
}

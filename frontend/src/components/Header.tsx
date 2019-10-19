import React from "react"

export default class Header extends React.Component {
	constructor(props) {
		super(props)
		this.state = {
			name: "Guest",
		}
	}
	render() {
		return (
			<header>
				<h1>Todos App</h1>
				ようこそ {this.state.name}さん
			</header>
		)
	}
}

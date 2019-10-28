import React from "react"
import PropTypes from "prop-types"

export default class TaskList extends React.Component {
	static propTypes = {
		fetch: PropTypes.func.isRequired,
		table: PropTypes.object.isRequired,
	}

	constructor(props) {
		super(props)
	}
	componentDidMount() {
		this.props.fetch()
	}
	render() {
		return this.props.table
	}
}

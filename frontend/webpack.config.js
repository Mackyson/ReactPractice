module.exports = {
	entry: "./src/index.jsx", // エントリポイントのjsxファイル
	output: {
		path: `${__dirname}` + "/public",
		filename: "bundle.js", // 出力するファイル
	},
	module: {
		rules: [
			{
				test: /\.(j|t)s(x?)$/, //jsファイル，またはjsxファイル
				exclude: /node_modules/, // node_modulesフォルダ配下でなければ
				loader: "babel-loader", // babel-loaderを使って変換する
				query: {
					plugins: ["transform-react-jsx"], // babelのtransform-react-jsxプラグインを使ってjsxを変換
				},
			},
		],
	},
}

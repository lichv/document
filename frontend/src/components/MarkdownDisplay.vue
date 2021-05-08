
<template>
	<div class="markdown-display">
		<div class="markdown-header">

		</div>
		<div class="markdown-content">
			<div class="markdown-side">
				<el-menu>
					<el-menu-item index="1-1">选项1</el-menu-item>
				</el-menu>
			</div>
			<div class="markdown-main">
				
			</div>
		</div>
	</div>
</template>
<script>
	import api from '/@/api';
	export default {
		name: 'ElMarkdownDisplay',
		props: {
			stsUrl: {
				type: String,
				default: "/api/aliyun/sts"
			},
			region: {
				type: String,
				default: "oss-cn-shanghai"
			},
			bucket: {
				type: String,
				default: "easytc"
			},
			rootname: {
				type: String,
				default: "aegicare"
			},
			dirname: {
				type: String,
				default: "picture"
			},
		},
		data() {
			return {
				client:null,
				imageUrl: '',
				items:[],
			}
		},
		mounted() {
			this.getFiles()
		},
		methods:{
			
			getFiles:function(){
				let _this = this;
				api.getMarkdownFiles.send()
				.then(result => {
					console.log('result',result)
					if (result.state==2000) {
						_this.items = result.data
					}
				})
				.catch(e => {
					console.log(e)
				})
			}
		},
	}
</script>
<style>
.markdown-display{
	height: 100%;
	width: 100%;
	padding: 0;
	margin: 0;
}
.markdown-header{
	height: 50px;
	background: rgb(225, 243, 216);
}
.markdown-content{
	height: calc(100% - 50px);
	background: rgb(250, 236, 216);
}
.markdown-side{
	background: rgb(233, 233, 235);
	width: 240px;
	height: 100%;
}
</style>
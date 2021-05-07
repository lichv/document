
<template>
	<div class="markdown-display">
		<div class="markdown-header">
			
		</div>
		<div class="markdown-content">
			<div class="markdown-side">
				
			</div>
			<div class="markdown-main">
				
			</div>
		</div>
	</div>
</template>
<script>

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
			}
		},
		mounted() {

		},
		methods:{


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
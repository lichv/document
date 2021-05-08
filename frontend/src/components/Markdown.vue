
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
		name: 'ElMarkdown',
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
}
.markdown-content{
	height: calc(100% - 50px);
}
</style>
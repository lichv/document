<template>
	<div class="oss_file">
		<input type="file" :id="id" :multiple="multiple" @change="doUpload"/>
	</div>
</template>
<script>
	import api from '/@/api';
	export default{
		name: 'ElUploadOss',
		data(){
			return{
				id:Math.random().toString(36).substr(2),
			};
		},
		props: {
			region: {
				type: String,
				default: "oss-cn-shanghai"
			},
			bucket: {
				type: String,
				default: "easyun"
			},
			rootname: {
				type: String,
				default: "aegicare"
			},
			dirname: {
				type: String,
				default: "picture"
			},
			multiple:{
				type: Boolean,
				twoWay:false
			},
		},
		
		methods:{
			doUpload(){
				const _this = this;
				api.getAliyunSTS.send().then((result) => {
					const client = new OSS({
						region:_this.region,
						accessKeyId: result.data.AccessKeyId,
						accessKeySecret: result.data.AccessKeySecret,
						stsToken: result.data.SecurityToken,
						bucket:_this.bucket
					})
					const files = document.getElementById(_this.id);
					if(files.files){
						const fileLen = document.getElementById(_this.id).files
						const resultUpload = [];		
						for (let i = 0; i < fileLen.length; i++) {
							const file = fileLen[i];
							const storeAs = file.name;
							var reader = new window.FileReader()
							reader.readAsArrayBuffer(file)
							client.put(storeAs, file, {

							}).then((results) => {
								if(results.url){
									resultUpload.push(results.url);
								}
							}).catch((err) => {
								console.log(err);
							});
						}
						_this.url = resultUpload;
					}   
				});
			}
		}
	};
</script>
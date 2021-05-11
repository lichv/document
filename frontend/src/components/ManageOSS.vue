<template>
	<el-card class="upload-oss-card">
		<template #header>
			<div class="card-header">
				<span>OSS文件管理</span>
				<el-upload class="upload-alioss" multiple :show-file-list="false" :http-request="fnUploadRequest" :on-success="handleUploadSuccess" :before-upload="beforeFileUpload">
					<el-button size="small" type="primary" icon="el-icon-upload">点击上传</el-button>
				</el-upload>
			</div>
		</template>
		<el-table :data="fileList" style="width: 100%">
			<el-table-column prop="url" label="预览" width="180">
				<template #default="scope">
					<el-popover placement="top-start" :title="scope.row.name" :width="'auto'" trigger="hover" >
						<template #reference>
							<img  v-if="scope.row.url" shape="square" :src="scope.row.url" style="max-height: 48px;">
						</template>
						<el-card>
							<img  v-if="scope.row.url" shape="square" :src="scope.row.url" style="max-width: 100%;">
						</el-card>
					</el-popover>
				</template>
			</el-table-column>
			<el-table-column prop="url" label="文件地址">
				<template #default="scope">
					<p class="copybtn" @click="copyShaneUrl(scope.row.url)" >{{scope.row.url}}</p>
				</template>
			</el-table-column>
		</el-table>
	</el-card>
</template>
<script>
	import api from '/@/api';
	export default {
		name: 'ElManageOss',
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
		},
		data() {
			return {
				client:null,
				imageUrl: '',
				maxKeys:2,
				marker:null,
				fileList: [],
			}
		},
		mounted() {
			const _this = this;
			api.getAliyunSTS.send()
			.then(result => {
				if (result.state==2000) {
					_this.client = new OSS({
						region:_this.region,
						accessKeyId: result.data.AccessKeyId,
						accessKeySecret: result.data.AccessKeySecret,
						stsToken: result.data.SecurityToken,
						bucket:_this.bucket
					})
					_this.getOssFileList()
				}
			})
			.catch(e => {
				console.log(e)
			})
		},
		methods:{
			handleUploadSuccess(res) {
				if (res) this.imageUrl = res.url
			},
		beforeFileUpload(file) {
			const isPicture = file.type === 'image/jpeg' || file.type === 'image/png' || file.type === 'image/gif' ;
			const isLt20M = file.size / 1024 / 1024 < 20;

			if (!isPicture) {
				this.$message.error('之恩呢上传JPEG,PNG,GIF格式图片!');
			}
			if (!isLt20M) {
				this.$message.error('上传头像图片大小不能超过 2MB!');
			}
			return isPicture && isLt20M;
		},
		async fnUploadRequest(options) {
			const _this = this;
			try {
				let file = options.file;
				let date = new Date().getTime()
				let fileNames = `${date}_${file.name}`
				this.client.put(this.rootname+'/'+this.dirname+'/'+fileNames, file).then(res=>{
					if (res.res.statusCode === 200) {
						options.onSuccess(res)
						_this.$emit('uploadSuccess',res)
						_this.getOssFileList()
					}else {
						options.onError("上传失败")
					}
				})
			}catch (e) {
				console.log('error',e)
				options.onError("上传失败")
			}
		},
		async getOssFileList(){
			const _this = this;
			let fileList = []
			if (this.client) {
				do {
					let result = await _this.client.list({
						'prefix':_this.rootname+'/'+_this.dirname+'/',
						'marker': _this.marker,
						'max-keys': _this.maxKeys
					});
					_this.marker = result.nextMarker;
					if (result.res.statusCode==200) {
						if (result.objects) {
							for (var i = result.objects.length - 1; i >= 0; i--) {
								var temp = result.objects[i]
								temp['name'] = temp['name'].substr(temp['name'].lastIndexOf('/')+1)
								fileList.push(temp)
							}
							
						}
					}
				}while(_this.marker)
				_this.fileList = fileList
			}
			console.log('fileList',fileList)
		},
		copyShaneUrl(shareLink){
			var _input = document.createElement("input");
			_input.value = shareLink;
			document.body.appendChild(_input);
			_input.select();
			document.execCommand("Copy");
			document.body.removeChild(_input);
			this.$message({
				showClose: true,
				message: "复制成功",
				type: "success"
			});
		},
	},
}
</script>
<style>
.upload-oss-card{
	height: 100%;
}
.upload-oss-card.el-card > .el-card__body{
	height: calc(100% - 80px);
	overflow-y: scroll;
}
.card-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-right: 50px;
    height: 50px;
}
.el-popover.el-popper{
	padding: 4px;
}
</style>
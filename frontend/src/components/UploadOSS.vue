<template>
	<div class="UploadOSS">
		<el-upload class="upload-alioss" action="" :show-file-list="false" :http-request="fnUploadRequest" :on-success="handleUploadSuccess" :before-upload="beforeAvatarUpload">
			<el-button size="small" type="primary">点击上传</el-button>
		</el-upload>
	</div>
</template>
<script>
import OSS from 'ali-oss'
import api from '/@/api'
export default {
	name: 'ElManageOSS',
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
		const _this = this;
		api.getAliyunSTS.send()
		.then(result => {
			console.log(result)
			if (result.state==2000) {
				_this.client = new OSS({
					region:_this.region,
					accessKeyId: result.data.Credentials.AccessKeyId,
					accessKeySecret: result.data.Credentials.AccessKeySecret,
					stsToken: result.data.Credentials.SecurityToken,
					bucket:_this.bucket
				})
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
		beforeAvatarUpload(file) {
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
			try {
				let file = options.file;
				let fileNames = file.name
				console.log('fileNames',fileNames)
				this.client.put(this.rootname+'/'+this.dirname+'/'+fileNames, file).then(res=>{
					if (res.res.statusCode === 200) {
						console.log('success',res)
						options.onSuccess(res)
						this.$emit('uploadSuccess',res)
					}else {
						console.log('failed')
						options.onError("上传失败")
					}
				})
			}catch (e) {
				console.log('error',e)
				options.onError("上传失败")
			}
		},

	},
}
</script>
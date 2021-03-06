# document

## 后端

```shell
python -V
python -m venv venv
venv\Scripts\activate
pip install -r requirements.txt
python web.py
```

## 前端

安装node.js

```shell
npm install
npm run dev
npm run build
```

创建项目

```shell
npm install -g create-vite-app
create-vite-app frontend
cd frontend
npm install element-plus -S
npm install vue-router -S
```

创建 vite.config.js 文件

```js
const path = require("path");
// vite.config.js # or vite.config.ts
console.log(path.resolve(__dirname, "./src"));

module.exports = {
  alias: {
    // 键必须以斜线开始和结束
    "/@/": path.resolve(__dirname, "./src"),
  },
  /**
   * 在生产中服务时的基本公共路径。
   * @default '/'
   */
  base: "./",
  /**
   * 与“根”相关的目录，构建输出将放在其中。如果目录存在，它将在构建之前被删除。
   * @default 'dist'
   */
  outDir: "dist",
  port: 3000,
  // 是否自动在浏览器打开
  open: true,
  // 是否开启 https
  https: false,
  // 服务端渲染
  ssr: false,
  // 引入第三方的配置
  //   optimizeDeps: {
  //     include: ["moment", "echarts", "axios", "mockjs"],
  //   },
  proxy: {
    // 如果是 /api 打头，则访问地址如下
    "/api": {
      target: "https://baidu.com/",
      changeOrigin: true,
      rewrite: (path) => path.replace(/^\/api/, ""),
    },
  },
};

```

启动项目

```shell
npm install
npm run dev
```


# Wayne Frontend

Wayne 的前端工程，基于 Angular 8 和 Clarity 构建，包含普通用户入口 `portal` 和管理员入口 `admin`。

## 技术栈

- Angular CLI 8.1
- Angular 8.1
- TypeScript 3.4
- Clarity UI 2.x
- RxJS 6

## 目录结构

```text
frontend/
+-- src/
|   +-- app/
|   |   +-- admin/      # 管理员端页面与路由
|   |   +-- portal/     # 普通用户端页面与路由
|   |   +-- shared/     # 公共组件、服务、模型与工具
|   +-- assets/         # 静态资源与国际化文案
|   +-- component/      # 全局样式组件
|   +-- environments/   # 环境配置
+-- lib/                # Wayne 前端插件代码
+-- deploy/             # 部署相关文件
+-- e2e/                # 端到端测试
+-- proxy.config.js     # 本地开发代理配置
+-- angular.json        # Angular CLI 配置
```

## 环境要求

- Node.js：建议使用 10.x 或 12.x
- npm：建议使用随 Node.js 安装的版本
- 后端 API：本地开发默认代理到 `http://localhost:8080`

> 如果依赖安装失败，可以先切换到国内 npm 镜像源：
>
> ```bash
> npm config set registry https://mirrors.huaweicloud.com/repository/npm/
> ```

## 本地开发

1. 初始化子模块：

   ```bash
   git submodule update --init --recursive
   ```

2. 启动 Wayne 后端服务，确保 `http://localhost:8080` 可访问。

3. 安装依赖：

   ```bash
   npm install
   ```

4. 启动开发服务：

   ```bash
   npm run start
   ```

5. 打开浏览器访问：

   ```text
   http://localhost:4200
   ```

开发服务会启用 HMR。接口请求会由 `src/config.js` 中的 `window.CONFIG.URL` 指向后端，默认地址为 `http://localhost:8080`。

## 常用命令

```bash
npm run start      # 启动本地开发环境
npm run build      # 构建开发包，输出到 dist/
npm run build:aot  # 生产模式 AOT 构建
npm run test       # 运行单元测试
npm run lint       # 运行 TSLint 检查
npm run e2e        # 运行端到端测试
```

## 构建

执行：

```bash
npm run build
```

构建产物会输出到 `dist/` 目录。

生产构建执行：

```bash
npm run build:aot
```

## 配置说明

- `src/config.js`：运行时配置，默认 API 地址为 `http://localhost:8080`。
- `proxy.config.js`：本地开发代理配置，默认目标为 `http://localhost:8080`。当前 `npm run start` 未显式启用该代理，如需使用可在启动命令中追加 `--proxy-config proxy.config.js`。
- `src/environments/`：Angular 编译时环境配置。

## 插件开发

插件代码位于 `lib/`：

- `lib/shared`：插件公共 model、client 与工具代码。
- `lib/portal`：普通用户端插件页面。
- `lib/admin`：管理员端插件页面。

新增插件时，通常需要在 `shared` 中补充 model/client，在 `portal` 或 `admin` 中新增组件，并将模块和路由注册到对应的 `library-*.module.ts` 与 `library-routing-*.ts`。

## 常见问题

### 依赖安装异常

可以删除 `node_modules` 后重新安装：

```bash
rm -rf node_modules
npm install
```

Windows PowerShell：

```powershell
Remove-Item -Recurse -Force node_modules
npm install
```

### 接口请求失败

确认后端服务已经启动，并检查 `src/config.js` 中的 `window.CONFIG.URL` 是否指向正确的后端地址。

import Link from 'next/link';
import { ThemeToggle } from '@/components/ThemeToggle';

export default function Home() {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      {/* 导航栏 */}
      <nav className="border-b border-gray-200 dark:border-gray-800 bg-white/90 dark:bg-gray-900/90 backdrop-blur supports-[backdrop-filter]:bg-white/60 dark:supports-[backdrop-filter]:bg-gray-900/60">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <span className="text-xl font-semibold text-gray-800 dark:text-gray-100">
                Modern Stack
              </span>
            </div>
            <div className="flex items-center space-x-4">
              <ThemeToggle />
              <Link
                href="/login"
                className="inline-flex items-center px-4 py-2 text-sm font-medium text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors"
              >
                登录
              </Link>
              <Link
                href="/register"
                className="inline-flex items-center px-4 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-colors"
              >
                注册
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* 主要内容 */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16 sm:py-24">
        {/* 头部介绍 */}
        <div className="text-center space-y-8 mb-16">
          <h1 className="text-4xl sm:text-5xl font-bold tracking-tight text-gray-800 dark:text-gray-100">
            现代全栈应用开发模板
          </h1>
          <p className="max-w-2xl mx-auto text-xl text-gray-600 dark:text-gray-400">
            基于 Go 和 Next.js 的全栈应用开发模板，集成最新技术栈，开箱即用。
          </p>
          <div className="flex justify-center gap-4">
            <Link
              href="/register"
              className="inline-flex items-center px-6 py-3 text-base font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-colors"
            >
              开始使用
              <svg
                className="ml-2 h-4 w-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth={2}
                  d="M14 5l7 7m0 0l-7 7m7-7H3"
                />
              </svg>
            </Link>
            <a
              href="https://github.com/your-username/modern-stack"
              className="inline-flex items-center px-6 py-3 text-base font-medium text-gray-600 dark:text-gray-300 bg-white dark:bg-gray-800 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg transition-colors border border-gray-200 dark:border-gray-700"
              target="_blank"
              rel="noopener noreferrer"
            >
              查看文档
            </a>
          </div>
        </div>

        {/* 特性列表 */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          {/* 前端特性 */}
          <div className="p-6 bg-white dark:bg-gray-800 rounded-2xl border border-gray-200 dark:border-gray-700">
            <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-100 mb-4">
              现代前端技术栈
            </h2>
            <ul className="space-y-3 text-gray-600 dark:text-gray-400">
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                Next.js 15 App Router
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                React 19 新特性
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                TypeScript 5
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                TailwindCSS 4
              </li>
            </ul>
          </div>

          {/* 后端特性 */}
          <div className="p-6 bg-white dark:bg-gray-800 rounded-2xl border border-gray-200 dark:border-gray-700">
            <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-100 mb-4">
              高性能后端架构
            </h2>
            <ul className="space-y-3 text-gray-600 dark:text-gray-400">
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                Go 1.22
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                Gin Web 框架
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                GORM
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                PostgreSQL
              </li>
            </ul>
          </div>

          {/* 开发体验 */}
          <div className="p-6 bg-white dark:bg-gray-800 rounded-2xl border border-gray-200 dark:border-gray-700">
            <h2 className="text-lg font-semibold text-gray-800 dark:text-gray-100 mb-4">
              优秀开发体验
            </h2>
            <ul className="space-y-3 text-gray-600 dark:text-gray-400">
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                Docker 容器化
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                JWT 认证
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                API 文档
              </li>
              <li className="flex items-center">
                <span className="mr-3 text-blue-500">●</span>
                自动化测试
              </li>
            </ul>
          </div>
        </div>
      </main>

      {/* 页脚 */}
      <footer className="border-t border-gray-200 dark:border-gray-800">
        <div className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
          <p className="text-center text-gray-500 dark:text-gray-400">
            © 2024 Modern Stack. All rights reserved.
          </p>
        </div>
      </footer>
    </div>
  );
}

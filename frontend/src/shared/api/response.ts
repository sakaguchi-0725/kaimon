import type { components } from './schema'

// OpenAPI Schemaから特定のスキーマ型を取得するユーティリティ型
export type ApiSchema<T extends keyof components['schemas']> = components['schemas'][T]

// レスポンス型のキー
type KeyofResponses = keyof components['responses']

// OpenAPI Schemaから特定のレスポンス型を取得するユーティリティ型
export type ApiResponse<T extends KeyofResponses> = components['responses'][T] extends {
  content: { 'application/json': unknown }
}
  ? components['responses'][T]['content']['application/json']
  : never

// OpenAPI Schemaから特定のクエリパラメータ型を取得するユーティリティ型
export type ApiQuery<T extends keyof components['parameters']> =
  components['parameters'][T] extends {
    schema: unknown
  }
    ? components['parameters'][T]['schema']
    : never

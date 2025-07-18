/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
  '/signup': {
    parameters: {
      query?: never
      header?: never
      path?: never
      cookie?: never
    }
    get?: never
    put?: never
    /** ユーザー登録 */
    post: {
      parameters: {
        query?: never
        header?: never
        path?: never
        cookie?: never
      }
      requestBody: {
        content: {
          'application/json': components['schemas']['SignUpRequest']
        }
      }
      responses: {
        /** @description ユーザー登録成功 */
        204: {
          headers: {
            [name: string]: unknown
          }
          content?: never
        }
        400: components['responses']['BadRequest']
        401: components['responses']['Unauthorized']
        500: components['responses']['InternalServerError']
      }
    }
    delete?: never
    options?: never
    head?: never
    patch?: never
    trace?: never
  }
  '/account': {
    parameters: {
      query?: never
      header?: never
      path?: never
      cookie?: never
    }
    /**
     * アカウント情報取得
     * @description 認証されたユーザーのアカウント情報を取得します
     */
    get: {
      parameters: {
        query?: never
        header?: never
        path?: never
        cookie?: never
      }
      requestBody?: never
      responses: {
        /** @description アカウント情報取得成功 */
        200: {
          headers: {
            [name: string]: unknown
          }
          content: {
            'application/json': components['schemas']['GetAccountResponse']
          }
        }
        401: components['responses']['Unauthorized']
        404: components['responses']['NotFound']
        500: components['responses']['InternalServerError']
      }
    }
    put?: never
    post?: never
    delete?: never
    options?: never
    head?: never
    patch?: never
    trace?: never
  }
  '/groups': {
    parameters: {
      query?: never
      header?: never
      path?: never
      cookie?: never
    }
    /**
     * 参加しているグループ一覧取得
     * @description ユーザーが参加しているグループの一覧を取得します
     */
    get: {
      parameters: {
        query?: never
        header?: never
        path?: never
        cookie?: never
      }
      requestBody?: never
      responses: {
        /** @description グループ一覧取得成功 */
        200: {
          headers: {
            [name: string]: unknown
          }
          content: {
            'application/json': components['schemas']['GetJoinedGroupsResponse']
          }
        }
        401: components['responses']['Unauthorized']
        500: components['responses']['InternalServerError']
      }
    }
    put?: never
    /**
     * グループ作成
     * @description 新しいグループを作成します
     */
    post: {
      parameters: {
        query?: never
        header?: never
        path?: never
        cookie?: never
      }
      requestBody: {
        content: {
          'application/json': components['schemas']['CreateGroupRequest']
        }
      }
      responses: {
        /** @description グループ作成成功 */
        201: {
          headers: {
            [name: string]: unknown
          }
          content?: never
        }
        400: components['responses']['BadRequest']
        401: components['responses']['Unauthorized']
        422: components['responses']['GroupLimitExceeded']
        500: components['responses']['InternalServerError']
      }
    }
    delete?: never
    options?: never
    head?: never
    patch?: never
    trace?: never
  }
  '/groups/{id}': {
    parameters: {
      query?: never
      header?: never
      path?: never
      cookie?: never
    }
    /**
     * グループ詳細取得
     * @description 指定されたグループの詳細情報を取得します
     */
    get: {
      parameters: {
        query?: never
        header?: never
        path: {
          /** @description グループID */
          id: string
        }
        cookie?: never
      }
      requestBody?: never
      responses: {
        /** @description グループ詳細取得成功 */
        200: {
          headers: {
            [name: string]: unknown
          }
          content: {
            'application/json': components['schemas']['GetGroupResponse']
          }
        }
        401: components['responses']['Unauthorized']
        403: components['responses']['Forbidden']
        404: components['responses']['NotFound']
        500: components['responses']['InternalServerError']
      }
    }
    put?: never
    post?: never
    delete?: never
    options?: never
    head?: never
    patch?: never
    trace?: never
  }
  '/groups/{id}/items': {
    parameters: {
      query?: never
      header?: never
      path?: never
      cookie?: never
    }
    /**
     * ショッピングアイテム一覧取得
     * @description 指定されたグループのショッピングアイテム一覧を取得します
     */
    get: {
      parameters: {
        query?: {
          /** @description カテゴリーID */
          categoryId?: number
          /** @description ステータス */
          status?: components['schemas']['ShoppingItemStatus']
          /** @description 取得件数（デフォルト：20） */
          limit?: number
          /** @description 取得開始位置（デフォルト：0） */
          offset?: number
        }
        header?: never
        path: {
          /** @description グループID */
          id: string
        }
        cookie?: never
      }
      requestBody?: never
      responses: {
        /** @description ショッピングアイテム一覧取得成功 */
        200: {
          headers: {
            [name: string]: unknown
          }
          content: {
            'application/json': components['schemas']['GetShoppingItemsResponse']
          }
        }
        401: components['responses']['Unauthorized']
        403: components['responses']['Forbidden']
        404: components['responses']['NotFound']
        500: components['responses']['InternalServerError']
      }
    }
    put?: never
    post?: never
    delete?: never
    options?: never
    head?: never
    patch?: never
    trace?: never
  }
}
export type webhooks = Record<string, never>
export interface components {
  schemas: {
    SignUpRequest: {
      /** @description ユーザー名 */
      name: string
      /** @description プロフィール画像URL */
      profileImageUrl?: string
    }
    GetJoinedGroupsResponse: {
      groups: components['schemas']['JoinedGroup'][]
    }
    JoinedGroup: {
      /** @description グループID */
      id: string
      /** @description グループ名 */
      name: string
      /** @description メンバー数 */
      memberCount: number
    }
    CreateGroupRequest: {
      /** @description グループ名 */
      name: string
      /** @description グループ説明 */
      description?: string
    }
    Error: {
      /** @description エラーコード */
      code: string
      /** @description エラーメッセージ */
      message: string
    }
    GetGroupResponse: {
      /** @description グループID */
      id: string
      /** @description グループ名 */
      name: string
      /** @description グループ説明 */
      description: string
      /**
       * Format: date-time
       * @description グループ作成日時
       */
      createdAt: string
      /** @description グループメンバー一覧 */
      members: components['schemas']['Member'][]
    }
    Member: {
      /** @description メンバーID */
      id: string
      /** @description メンバー名 */
      name: string
      /** @description メンバーの役割 */
      role: components['schemas']['MemberRole']
      /** @description メンバーのステータス */
      status: components['schemas']['MemberStatus']
    }
    /**
     * @description メンバーの役割
     * @enum {string}
     */
    MemberRole: 'admin' | 'member'
    /**
     * @description メンバーのステータス
     * @enum {string}
     */
    MemberStatus: 'active' | 'pending'
    GetShoppingItemsResponse: {
      items: components['schemas']['ShoppingItem'][]
      /**
       * Format: int64
       * @description 全アイテム数
       */
      totalCount: number
      /** @description 次のページがあるかどうか */
      hasNext: boolean
    }
    ShoppingItem: {
      /**
       * Format: int64
       * @description アイテムID
       */
      id: number
      /** @description アイテム名 */
      name: string
      /** @description 登録者ID */
      memberId: string
      status: components['schemas']['ShoppingItemStatus']
    }
    /**
     * @description 買い物メモのステータス
     * @enum {string}
     */
    ShoppingItemStatus: 'UNPURCHASED' | 'PURCHASED' | 'IN_CART'
    GetAccountResponse: {
      /** @description アカウントID */
      id: string
      /** @description アカウント名 */
      name: string
    }
  }
  responses: {
    /** @description リクエスト不正 */
    BadRequest: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "BAD_REQUEST",
         *       "message": "リクエストが不正です"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description 確認コードの有効期限切れ */
    ExpiredCode: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "EXPIRED_CODE",
         *       "message": "確認コードの有効期限が切れています"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description 認証エラー */
    Unauthorized: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "UNAUTHORIZED",
         *       "message": "メールアドレスまたはパスワードが正しくありません"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description アクセス権限がありません */
    Forbidden: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "FORBIDDEN",
         *       "message": "リソースへのアクセス権限がありません"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description リソースが見つかりません */
    NotFound: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "NOT_FOUND",
         *       "message": "グループが見つかりません"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description グループ数制限超過 */
    GroupLimitExceeded: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "GROUP_LIMIT_EXCEEDED",
         *       "message": "グループ数の上限に達しています。プランをアップグレードすると、より多くのグループを作成できます。"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
    /** @description サーバーエラー */
    InternalServerError: {
      headers: {
        [name: string]: unknown
      }
      content: {
        /** @example {
         *       "code": "INTERNAL_SERVER_ERROR",
         *       "message": "サーバーエラーが発生しました"
         *     } */
        'application/json': components['schemas']['Error']
      }
    }
  }
  parameters: never
  requestBodies: never
  headers: never
  pathItems: never
}
export type $defs = Record<string, never>
export type operations = Record<string, never>

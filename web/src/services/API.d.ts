declare namespace API {
  export interface PageQuery {
    filters?: Map<string, any>;
    pageNo?: number;                                                                                      `json:"pageNo"`
    pageSize?: number                                                                            `json:"pageSize"`
    sort?:SortSpec[]
  }

  export interface SortSpec {
    property?: string
    type?: 'ASC' | 'DSC'
  }

  export interface Page<T> {
    total?: number;
    pageNo?: number;
    pageSize?: number;
    pageCount?: number;
    items?: T[];
  }

  export interface CurrentUser {
    avatar?: string;
    name?: string;
    title?: string;
    group?: string;
    signature?: string;
    tags?: {
      key: string;
      label: string;
    }[];
    userid?: string;
    access?: 'user' | 'guest' | 'admin';
    unreadCount?: number;
  }

  export interface LoginStateType {
    status?: 'ok' | 'error';
    type?: string;
  }

  export interface NoticeIconData {
    id: string;
    key: string;
    avatar: string;
    title: string;
    datetime: string;
    type: string;
    read?: boolean;
    description: string;
    clickClose?: boolean;
    extra: any;
    status: string;
  }

  export interface CustomObject {
    objId?: string;
    orgId?: string;
    objName?: string;
  }

  export interface CustomField {
    fieldId?: string;
    objId?: string;
    orgId?: string;
    defaultValue?: string;
    deprecated?: boolean;
    description?: string;
    displayFormat?: string;
  }
}

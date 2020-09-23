import { request } from 'umi';

export async function page(umi: API.PageQuery) {
  return request<API.Page<CustomObject>>('/api/custom_objects/page', {
    method: 'POST',
    data: params,
  });
}

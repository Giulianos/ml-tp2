import { api } from '@/config/api'

export const create = (csv, predAttr, gainFunc, minNodeCount) => api.post(`/tree?predAttr=${predAttr}&gainFunc=${gainFunc}&minNodeCount=${minNodeCount}`, csv)

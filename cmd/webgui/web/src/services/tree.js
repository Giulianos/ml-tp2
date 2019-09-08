import { api } from '@/config/api'

export const create = (csv, predAttr) => api.post(`/tree?pred-attr=${predAttr}`, csv)

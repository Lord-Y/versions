export interface Loading {
  loading: boolean
  class: string
}

export interface Pagination {
  url: string
  actualPage: number
  total: number
  mode?: string
}

export interface Deployment {
  changelog_url: string
  date: Date
  environment: string
  platform: string
  raw: string
  status: string
  version: string
  versions_id: number
  workload: string
  total: number
}

export type Deployments = Array<Deployment>

export interface GenericObject {
  [key: string]: string | number | boolean
}

export interface Workload {
  workload: string
  platform: string
  environment: string
}

export type Workloads = Array<Workload>

export interface StatLatest {
  total: number
  workload: string
  platform: string
  environment: string
  status: string
  date: Date
}

export type StatsLatest = Array<StatLatest>

export interface BarData {
  labels: Array<string>
  datasets: dataset[]
}

export interface dataset {
  label: string
  data: Array<number>
  backgroundColor: Array<string>
  // borderColor: string
}

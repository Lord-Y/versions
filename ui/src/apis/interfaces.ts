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

export interface DeploymentRaw {
  raw: string
}

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

export interface dataset {
  label: string
  data: Array<number>
  backgroundColor: Array<string>
}

export interface BarData {
  labels: string[]
  datasets: dataset[]
}

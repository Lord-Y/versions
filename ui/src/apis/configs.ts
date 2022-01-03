export interface configObject {
  RANGE_LIMIT: number
}

export const config = {
  RANGE_LIMIT: process.env['RANGE_LIMIT'] || 25,
} as configObject

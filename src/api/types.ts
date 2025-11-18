export interface HueAreYouRecord {
  name: string
  choice: Record<string, string>
}

export interface SessionData {
  user_id: string
  token: string
}

export interface AdminLoginPayload {
  name: string
  password: string
}

export interface AdminSignInPayload {
  name: string
  email: string
  password: string
}

export interface SaveHueAreYouResultPayload {
  user_name: string
  record: HueAreYouRecord
}

export interface FetchHueAreYouDataParams {
  session: SessionData
  dataRange: [number, number]
}

export interface HueAreYouDataResponse {
  records: HueAreYouRecord[]
}

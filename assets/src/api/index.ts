import axios from 'axios'

export const backendURL = import.meta.env.VITE_BACKEND ?? `${window.location.protocol}//${window.location.host}`

export const backendAPIURL = backendURL + '/api'

export const backend = axios.create({
    baseURL: backendAPIURL
})
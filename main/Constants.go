package main

import "time"

// PageSize is default number of records for paging
const PageSize = 20

// ApplicationID - app ID
const ApplicationID = "8ef0474a-4947-4c35-b171-900b026ae30f"

// AppSecret - app secret
const AppSecret = "TaskManagementSecret2018"

// TokenValidPeriodInMinutes - period which token is valid from the time it's issued
const TokenValidPeriodInMinutes = 120

// ClaimAppID - app id
const ClaimAppID = "AppID"

// ClaimUserID - user ID
const ClaimUserID = "UserID"

// ClaimEmail - email
const ClaimEmail = "Email"

// ClaimExpiredTime - expired time
const ClaimExpiredTime = "ExpiredTime"

// ClaimTimeFormat - claim time format
const ClaimTimeFormat = time.RFC3339

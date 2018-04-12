package main

import "time"

// PageSize is default number of records for paging
const PageSize = 20

// UserStatusInactive - represent inactive status of user
const UserStatusInactive = 0

// UserStatusActive - represent active status of user
const UserStatusActive = 1

// UserStatusDeleted - represent deleted status of user
const UserStatusDeleted = 2

// ApplicationID - app ID
const ApplicationID = "8ef0474a-4947-4c35-b171-900b026ae30f"

// AppSecret - app secret
const AppSecret = "TaskManagementSecret2018"

// TokenValidPeriodInMinutes - period which token is valid from the time it's issued
const TokenValidPeriodInMinutes = 120

// ClaimAppID - app id
const ClaimAppID = "appID"

// ClaimEmail - email
const ClaimEmail = "email"

// ClaimExpiredTime - expired time
const ClaimExpiredTime = "expiredTime"

// ClaimTimeFormat - claim time format
const ClaimTimeFormat = time.RFC3339

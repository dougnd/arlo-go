package arlo

const (
	BaseUrl               = "https://arlo.netgear.com/hmsweb"
	LoginUri              = "/login/v2"
	LogoutUri             = "/logout"
	SubscribeUri          = "/client/subscribe?token=%s"
	UnsubscribeUri        = "/client/unsubscribe"
	NotifyUri             = "/users/devices/notify/%s"
	ResetUri              = "/users/library/reset"
	ServiceLevelUri       = "/users/serviceLevel"
	OffersUri             = "/users/payment/offers"
	UserProfileUri        = "/users/profile"
	UserChangePasswordUri = "/users/changePassword"
	UserSessionUri        = "/users/session"
	UserFriendsUri        = "/users/friends"
	UserLocationsUri      = "/users/locations"
	UserLocationUri       = "/users/locations/%s"
	LibraryUri            = "/users/library"
	LibraryRecycleUri     = "/users/library/recycle"
	LibraryMetadataUri    = "/users/library/metadata"
	DevicesUri            = "/users/devices"
	DeviceRenameUri       = "/users/devices/renameDevice"
	DeviceDisplayOrderUri = "/users/devices/displayOrder"
	DeviceTakeSnapshotUri = "/users/devices/takeSnapshot"
	DeviceStartRecordUri  = "/users/devices/startRecord"
	DeviceStopRecordUri   = "/users/devices/stopRecord"
	DeviceStartStreamUri  = "/users/devices/startStream"
)

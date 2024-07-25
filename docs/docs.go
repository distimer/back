// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/logout": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "parameters": [
                    {
                        "description": "logoutTokenReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authctrl.logoutTokenReq"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/auth/oauth/apple": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Apple Oauth Login",
                "parameters": [
                    {
                        "description": "oauthLoginReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authctrl.oauthLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authctrl.loginRes"
                        }
                    }
                }
            }
        },
        "/auth/oauth/google": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Google Oauth Login",
                "parameters": [
                    {
                        "description": "oauthLoginReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authctrl.oauthLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authctrl.loginRes"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "description": "refreshTokenReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authctrl.refreshTokenReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authctrl.refreshTokenRes"
                        }
                    }
                }
            }
        },
        "/group": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Get All Joined Groups",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groupctrl.getJoinedGroupsRes"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Create Group",
                "parameters": [
                    {
                        "description": "createGroupReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/groupctrl.createGroupReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/group/invite/{id}": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Invite to Group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groupctrl.inviteGroupRes"
                        }
                    }
                }
            }
        },
        "/group/member/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Get All Group Members",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groupctrl.getAllGroupMembersRes"
                        }
                    }
                }
            }
        },
        "/group/policy/{id}": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Modify Group Policy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "modifyGroupPolicyReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/groupctrl.modifyGroupPolicyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/groupctrl.modifyGroupPolicyRes"
                        }
                    }
                }
            }
        },
        "/group/{id}": {
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Group"
                ],
                "summary": "Delete Group",
                "parameters": [
                    {
                        "type": "string",
                        "description": "group id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get My User Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userctrl.myUserInfoRes"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Modify User Info",
                "parameters": [
                    {
                        "description": "modifyUserInfoReq",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userctrl.modifyUserInfoReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userctrl.modifyUserInfoRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authctrl.loginRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "authctrl.logoutTokenReq": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "authctrl.oauthLoginReq": {
            "type": "object",
            "required": [
                "token"
            ],
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "authctrl.refreshTokenReq": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "authctrl.refreshTokenRes": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "ent.Affiliation": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the AffiliationQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.AffiliationEdges"
                        }
                    ]
                },
                "group_id": {
                    "description": "GroupID holds the value of the \"group_id\" field.",
                    "type": "string"
                },
                "joined_at": {
                    "description": "JoinedAt holds the value of the \"joined_at\" field.",
                    "type": "string"
                },
                "nickname": {
                    "description": "Nickname holds the value of the \"nickname\" field.",
                    "type": "string"
                },
                "role": {
                    "description": "Role holds the value of the \"role\" field.",
                    "type": "integer"
                },
                "user_id": {
                    "description": "UserID holds the value of the \"user_id\" field.",
                    "type": "string"
                }
            }
        },
        "ent.AffiliationEdges": {
            "type": "object",
            "properties": {
                "group": {
                    "description": "Group holds the value of the group edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Group"
                        }
                    ]
                },
                "user": {
                    "description": "User holds the value of the user edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                }
            }
        },
        "ent.Group": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "description": {
                    "description": "Description holds the value of the \"description\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the GroupQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.GroupEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "invite_policy": {
                    "description": "InvitePolicy holds the value of the \"invite_policy\" field.",
                    "type": "integer"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "nickname_policy": {
                    "description": "NicknamePolicy holds the value of the \"nickname_policy\" field.",
                    "type": "string"
                },
                "reveal_policy": {
                    "description": "RevealPolicy holds the value of the \"reveal_policy\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.GroupEdges": {
            "type": "object",
            "properties": {
                "invite_codes": {
                    "description": "InviteCodes holds the value of the invite_codes edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.InviteCode"
                    }
                },
                "members": {
                    "description": "Members holds the value of the members edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.User"
                    }
                },
                "owner": {
                    "description": "Owner holds the value of the owner edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                },
                "shared_study_logs": {
                    "description": "SharedStudyLogs holds the value of the shared_study_logs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.StudyLog"
                    }
                }
            }
        },
        "ent.InviteCode": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Code holds the value of the \"code\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the InviteCodeQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.InviteCodeEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "used": {
                    "description": "Used holds the value of the \"used\" field.",
                    "type": "boolean"
                }
            }
        },
        "ent.InviteCodeEdges": {
            "type": "object",
            "properties": {
                "group": {
                    "description": "Group holds the value of the group edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.Group"
                        }
                    ]
                }
            }
        },
        "ent.RefreshToken": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the RefreshTokenQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.RefreshTokenEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                }
            }
        },
        "ent.RefreshTokenEdges": {
            "type": "object",
            "properties": {
                "user": {
                    "description": "User holds the value of the user edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                }
            }
        },
        "ent.StudyLog": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "Content holds the value of the \"content\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the StudyLogQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.StudyLogEdges"
                        }
                    ]
                },
                "end_at": {
                    "description": "EndAt holds the value of the \"end_at\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "start_at": {
                    "description": "StartAt holds the value of the \"start_at\" field.",
                    "type": "string"
                }
            }
        },
        "ent.StudyLogEdges": {
            "type": "object",
            "properties": {
                "shared_group": {
                    "description": "SharedGroup holds the value of the shared_group edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Group"
                    }
                },
                "user": {
                    "description": "User holds the value of the user edge.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.User"
                        }
                    ]
                }
            }
        },
        "ent.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the UserQuery when eager-loading is set.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/ent.UserEdges"
                        }
                    ]
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "name": {
                    "description": "Name holds the value of the \"name\" field.",
                    "type": "string"
                },
                "oauth_id": {
                    "description": "OauthID holds the value of the \"oauth_id\" field.",
                    "type": "string"
                },
                "oauth_provider": {
                    "description": "OauthProvider holds the value of the \"oauth_provider\" field.",
                    "type": "integer"
                }
            }
        },
        "ent.UserEdges": {
            "type": "object",
            "properties": {
                "affilations": {
                    "description": "Affilations holds the value of the affilations edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Affiliation"
                    }
                },
                "joined_groups": {
                    "description": "JoinedGroups holds the value of the joined_groups edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Group"
                    }
                },
                "owned_groups": {
                    "description": "OwnedGroups holds the value of the owned_groups edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Group"
                    }
                },
                "refresh_tokens": {
                    "description": "RefreshTokens holds the value of the refresh_tokens edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.RefreshToken"
                    }
                },
                "study_logs": {
                    "description": "StudyLogs holds the value of the study_logs edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.StudyLog"
                    }
                }
            }
        },
        "groupctrl.createGroupReq": {
            "type": "object",
            "required": [
                "invite_policy",
                "name",
                "nickname",
                "reveal_policy"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "description between 0 and 100"
                },
                "invite_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                },
                "name": {
                    "type": "string",
                    "example": "name between 3 and 30"
                },
                "nickname": {
                    "type": "string",
                    "example": "nickname between 1 and 20"
                },
                "nickname_policy": {
                    "type": "string",
                    "example": "nickname_policy between 0 and 50"
                },
                "reveal_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                }
            }
        },
        "groupctrl.getAllGroupMembersRes": {
            "type": "object",
            "properties": {
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Affiliation"
                    }
                }
            }
        },
        "groupctrl.getJoinedGroupsRes": {
            "type": "object",
            "properties": {
                "joined_groups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Group"
                    }
                }
            }
        },
        "groupctrl.inviteGroupRes": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                }
            }
        },
        "groupctrl.modifyGroupPolicyReq": {
            "type": "object",
            "required": [
                "invite_policy",
                "reveal_policy"
            ],
            "properties": {
                "invite_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                },
                "reveal_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                }
            }
        },
        "groupctrl.modifyGroupPolicyRes": {
            "type": "object",
            "required": [
                "invite_policy",
                "reveal_policy"
            ],
            "properties": {
                "invite_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                },
                "reveal_policy": {
                    "type": "integer",
                    "maximum": 2,
                    "minimum": 0
                }
            }
        },
        "userctrl.modifyUserInfoReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "example": "name between 1 and 20"
                }
            }
        },
        "userctrl.modifyUserInfoRes": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "userctrl.myUserInfoRes": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Distimer Swagger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

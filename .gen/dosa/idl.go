// Code generated by thriftrw v1.12.0. DO NOT EDIT.
// @generated

package dosa

import "go.uber.org/thriftrw/thriftreflect"

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "dosa",
	Package:  "github.com/uber/dosa-idl/.gen/dosa",
	FilePath: "dosa.thrift",
	SHA1:     "ae00cb02a0e58de9a093e75609a9df13820eee71",
	Raw:      rawIDL,
}

const rawIDL = "// Copyright (c) 2017 Uber Technologies, Inc.\n\n// Permission is hereby granted, free of charge, to any person obtaining a copy\n// of this software and associated documentation files (the \"Software\"), to deal\n// in the Software without restriction, including without limitation the rights\n// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell\n// copies of the Software, and to permit persons to whom the Software is\n// furnished to do so, subject to the following conditions:\n\n// The above copyright notice and this permission notice shall be included in\n// all copies or substantial portions of the Software.\n\n// THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\n// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\n// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\n// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\n// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\n// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\n// THE SOFTWARE.\n\nnamespace java com.uber.dosa\n\ntypedef map<string, Value> FieldValueMap\n\nenum ElemType {\n   BOOL,\n   BLOB,\n   STRING,\n   INT32,\n   INT64,\n   DOUBLE,\n   TIMESTAMP,\n   UUID,\n\n   // Reserve a few enum types in case we want to support few more\n   RESERVED0,\n   RESERVED1,\n   RESERVED2,\n   RESERVED3,\n}\n\nunion RawValue {\n 1: optional binary binaryValue                 // BLOB, UUID\n 2: optional bool boolValue                     // BOOL\n 3: optional double doubleValue                 // DOUBLE\n 4: optional i32 int32Value                     // INT32\n 5: optional i64 (js.type = \"Long\") int64Value  // INT64, TIMESTAMP\n 6: optional string stringValue                 // STRING\n}\n\n// make it union in case we want to support collections like set\nunion Value {\n   1: optional RawValue elemValue\n}\n\nstruct SchemaRef {\n    1: optional string scope\n    2: optional string namePrefix\n    3: optional string entityName\n    4: optional i32 version\n}\n\nstruct FieldTag {\n   1: optional string name\n   2: optional string value\n}\n\nstruct FieldDesc {\n   1: optional ElemType type\n   2: optional set<FieldTag> tags\n}\n\nstruct ClusteringKey {\n   1: optional string name\n   2: optional bool asc\n}\n\nstruct PrimaryKey {\n   1: optional list<string> partitionKeys\n   2: optional list<ClusteringKey> clusteringKeys\n}\n\nstruct IndexDefinition {\n   1: optional PrimaryKey key\n}\n\nenum ETLState {\n    OFF = 1,\n    ON,\n    \n    // Reserve a few enum types in case we want to support more\n    RESERVED0,\n    RESERVED1\n}\n\nstruct EntityDefinition {\n   1: optional string name\n   2: optional map<string, FieldDesc> fieldDescs\n   3: optional PrimaryKey primaryKey\n   4: optional map<string, IndexDefinition> Indexes\n   5: optional ETLState etl\n}\n\nstruct Error {\n   1: optional i32 errCode\n   2: optional string msg\n   3: optional bool shouldRetry\n}\n\nstruct CreateRequest {\n   1: optional SchemaRef ref\n   2: optional FieldValueMap entityValues\n   3: optional i64 ttl\n}\n\nstruct ReadRequest {\n   1: optional SchemaRef ref\n   2: optional FieldValueMap keyValues\n   3: optional set<string> fieldsToRead\n}\n\nstruct ReadResponse {\n   1: optional FieldValueMap entityValues\n}\n\nstruct MultiReadRequest {\n   1: optional SchemaRef ref\n   2: optional list<FieldValueMap> keyValues\n   3: optional set<string> fieldsToRead\n}\n\nunion EntityOrError{\n   1: optional FieldValueMap entityValues\n   2: optional Error error\n}\n\nstruct MultiReadResponse {\n   1: optional list<EntityOrError> results\n}\n\nstruct MultiUpsertResponse {\n   1: optional list<Error> errors\n}\n\nstruct UpsertRequest {\n    1: optional SchemaRef ref\n    2: optional FieldValueMap entityValues\n    3: optional i64 ttl\n}\n\nstruct MultiUpsertRequest {\n    1: optional SchemaRef ref\n    2: optional list<FieldValueMap> entities\n    3: optional i64 ttl\n}\n\nstruct RemoveRequest {\n   1: optional SchemaRef ref\n   2: optional FieldValueMap keyValues\n}\n\nstruct MultiRemoveRequest {\n   1: optional SchemaRef ref \n   2: optional list<FieldValueMap> keyValues\n}\n\nstruct MultiRemoveResponse {\n   1: optional list<Error> errors\n}\n\nenum Operator {\n   EQ,\n   LT,\n   GT,\n   LT_OR_EQ,\n   GT_OR_EQ,\n}\n\nstruct Field {\n   1: optional string name\n   2: optional Value value\n}\n\nstruct Condition {\n   1: optional Operator op\n   2: optional Field field\n}\n\nstruct RangeRequest {\n   1: optional SchemaRef ref\n   2: optional string token\n   3: optional i32 limit\n   4: optional list<Condition> conditions\n   5: optional set<string> fieldsToRead\n}\n\nstruct RangeResponse {\n   1: optional list<FieldValueMap> entities\n   2: optional string nextToken\n}\n\nstruct RemoveRangeRequest {\n   1: optional SchemaRef ref\n   2: optional list<Condition> conditions\n}\n\nstruct SearchRequest {\n   1: optional SchemaRef ref\n   2: optional string token\n   3: optional i32 limit\n   4: optional Field searchBy\n   5: optional set<string> fieldsToRead\n}\n\nstruct SearchResponse {\n   1: optional list<FieldValueMap> entities\n   2: optional string nextToken\n}\n\nstruct ScanRequest {\n   1: optional SchemaRef ref\n   2: optional string token\n   3: optional i32 limit\n   4: optional set<string> fieldsToRead\n}\n\nstruct ScanResponse {\n   1: optional list<FieldValueMap> entities\n   2: optional string nextToken\n}\n\nstruct CanUpsertSchemaRequest {\n   1: optional string scope\n   2: optional string namePrefix\n   3: optional list<EntityDefinition> entityDefs\n}\n\nstruct CanUpsertSchemaResponse {\n   // latest upserted version\n   1: optional i32 version\n}\n\nstruct CheckSchemaRequest {\n   1: optional string scope\n   2: optional string namePrefix\n   3: optional list<EntityDefinition> entityDefs\n}\n\nstruct CheckSchemaResponse {\n   1: optional i32 version\n}\n\nstruct UpsertSchemaRequest {\n   1: optional string scope\n   2: optional string namePrefix\n   3: optional list<EntityDefinition> entityDefs\n   4: optional bool dryRun\n}\n\nstruct UpsertSchemaResponse {\n   1: optional i32 version\n   2: optional string status\n}\n\nstruct CheckSchemaStatusRequest {\n   1: optional string scope\n   2: optional string namePrefix\n   3: optional i32 version\n}\n\nstruct CheckSchemaStatusResponse {\n    1: optional i32 version\n    2: optional string status\n}\n\nstruct CreateScopeRequest {\n   1: optional string name\n}\n\nstruct ScopeExistsRequest {\n   1: optional string name\n}\n\nstruct ScopeExistsResponse {\n   1: optional bool exists\n}\n\nstruct TruncateScopeRequest {\n   1: optional string name\n}\nstruct DropScopeRequest {\n   1: optional string name\n}\n\nexception BadRequestError {\n    1: required string err\n    2: optional string message\n    3: optional i32 errorCode\n}\n\nexception InternalServerError {\n    1: required string err\n    2: optional string message\n    3: optional i32 errorCode\n}\n\nexception BadSchemaError {\n    // EntityName -> Error msg\n    1: required map<string, string> reasons\n}\n\nservice Dosa {\n   void createIfNotExists(\n       1: CreateRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   ReadResponse read (\n       1: ReadRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   MultiReadResponse multiRead (\n       1: MultiReadRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   void upsert (\n       1: UpsertRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   MultiUpsertResponse multiUpsert (\n       1: MultiUpsertRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n   \n   void remove (\n       1: RemoveRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   MultiRemoveResponse multiRemove (\n       1: MultiRemoveRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   RangeResponse range (\n       1: RangeRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   void removeRange (\n       1: RemoveRangeRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   SearchResponse search (\n       1: SearchRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   ScanResponse scan (\n       1: ScanRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   CanUpsertSchemaResponse canUpsertSchema(\n       1: CanUpsertSchemaRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n       3: BadSchemaError schemaError\n   )\n\n   CheckSchemaResponse checkSchema(\n       1: CheckSchemaRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n       3: BadSchemaError schemaError\n   )\n\n   UpsertSchemaResponse upsertSchema(\n       1: UpsertSchemaRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n       3: BadSchemaError schemaError\n   )\n\n   CheckSchemaStatusResponse checkSchemaStatus(\n       1: CheckSchemaStatusRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   void createScope(\n       1: CreateScopeRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   ScopeExistsResponse scopeExists(\n       1: ScopeExistsRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   void truncateScope(\n       1: TruncateScopeRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n\n   void dropScope(\n       1: DropScopeRequest request\n   ) throws (\n       1: BadRequestError clientError\n       2: InternalServerError serverError\n   )\n}\n\n"

//
//Hyperledger Cacti Plugin - Common Operator Primitives
//
//Contains/describes the Hyperledger Cacti Common Operator Primitives plugin.  These primitives require specific chaincode and Weaver relays to be deployed on the network as described at https://hyperledger-cacti.github.io/cacti/weaver/introduction/.
//
//The version of the OpenAPI document: 2.1.0
//
//Generated by OpenAPI Generator: https://openapi-generator.tech

// @generated by protoc-gen-es v1.8.0 with parameter "target=ts"
// @generated from file models/pledge_asset_v1200_response_pb.proto (package org.hyperledger.cacti.plugin.copm.core, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message org.hyperledger.cacti.plugin.copm.core.PledgeAssetV1200ResponsePB
 */
export class PledgeAssetV1200ResponsePB extends Message<PledgeAssetV1200ResponsePB> {
  /**
   * @generated from field: string pledge_id = 1;
   */
  pledgeId = "";

  constructor(data?: PartialMessage<PledgeAssetV1200ResponsePB>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "org.hyperledger.cacti.plugin.copm.core.PledgeAssetV1200ResponsePB";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pledge_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PledgeAssetV1200ResponsePB {
    return new PledgeAssetV1200ResponsePB().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PledgeAssetV1200ResponsePB {
    return new PledgeAssetV1200ResponsePB().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PledgeAssetV1200ResponsePB {
    return new PledgeAssetV1200ResponsePB().fromJsonString(jsonString, options);
  }

  static equals(a: PledgeAssetV1200ResponsePB | PlainMessage<PledgeAssetV1200ResponsePB> | undefined, b: PledgeAssetV1200ResponsePB | PlainMessage<PledgeAssetV1200ResponsePB> | undefined): boolean {
    return proto3.util.equals(PledgeAssetV1200ResponsePB, a, b);
  }
}


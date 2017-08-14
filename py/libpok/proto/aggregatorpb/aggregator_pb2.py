# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: aggregatorpb/aggregator.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='aggregatorpb/aggregator.proto',
  package='aggregatorpb',
  syntax='proto3',
  serialized_pb=_b('\n\x1d\x61ggregatorpb/aggregator.proto\x12\x0c\x61ggregatorpb\x1a-github.com/gogo/protobuf/gogoproto/gogo.proto\"\x82\x01\n\x0bHyperParams\x12\x1a\n\x12proportion_clients\x18\x01 \x01(\x01\x12\x12\n\nbatch_size\x18\x02 \x01(\x03\x12\x12\n\nnum_rounds\x18\x03 \x01(\x03\x12\x15\n\rlearning_rate\x18\x04 \x01(\x01\x12\x18\n\x10num_local_rounds\x18\x05 \x01(\x03\"\xaf\x01\n\x04Work\x12\'\n\x02id\x18\x01 \x01(\x0b\x32\x15.aggregatorpb.ModelIDB\x04\xc8\xde\x1f\x00\x12\x14\n\x0cnum_examples\x18\x02 \x01(\x03\x12\x13\n\x0bnum_clients\x18\x03 \x01(\x03\x12\r\n\x05\x65poch\x18\x04 \x01(\x03\x12\r\n\x05model\x18\x05 \x01(\x0c\x12\x35\n\x0chyper_params\x18\x07 \x01(\x0b\x32\x19.aggregatorpb.HyperParamsB\x04\xc8\xde\x1f\x00\"3\n\x07ModelID\x12\x0e\n\x06\x64omain\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\x03\":\n\x0eGetWorkRequest\x12(\n\x03ids\x18\x01 \x03(\x0b\x32\x15.aggregatorpb.ModelIDB\x04\xc8\xde\x1f\x00\";\n\x11ReportWorkRequest\x12&\n\x04work\x18\x01 \x03(\x0b\x32\x12.aggregatorpb.WorkB\x04\xc8\xde\x1f\x00\"\x11\n\x0fReportWorkReply\"<\n\x10ProdModelRequest\x12(\n\x03ids\x18\x01 \x03(\x0b\x32\x15.aggregatorpb.ModelIDB\x04\xc8\xde\x1f\x00\"\'\n\x11ProdModelResponse\x12\x12\n\nmodel_urls\x18\x02 \x03(\t*H\n\x0eTrainingStatus\x12\r\n\tSCHEDULED\x10\x00\x12\x0c\n\x08TRAINING\x10\x01\x12\r\n\tSUCCEEDED\x10\x02\x12\n\n\x06\x46\x41ILED\x10\x03\x32\xed\x01\n\nAggregator\x12?\n\x07GetWork\x12\x1c.aggregatorpb.GetWorkRequest\x1a\x12.aggregatorpb.Work\"\x00\x30\x01\x12N\n\nReportWork\x12\x1f.aggregatorpb.ReportWorkRequest\x1a\x1d.aggregatorpb.ReportWorkReply\"\x00\x12N\n\tProdModel\x12\x1e.aggregatorpb.ProdModelRequest\x1a\x1f.aggregatorpb.ProdModelResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[])

_TRAININGSTATUS = _descriptor.EnumDescriptor(
  name='TrainingStatus',
  full_name='aggregatorpb.TrainingStatus',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SCHEDULED', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TRAINING', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SUCCEEDED', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAILED', index=3, number=3,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=701,
  serialized_end=773,
)
_sym_db.RegisterEnumDescriptor(_TRAININGSTATUS)

TrainingStatus = enum_type_wrapper.EnumTypeWrapper(_TRAININGSTATUS)
SCHEDULED = 0
TRAINING = 1
SUCCEEDED = 2
FAILED = 3



_HYPERPARAMS = _descriptor.Descriptor(
  name='HyperParams',
  full_name='aggregatorpb.HyperParams',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='proportion_clients', full_name='aggregatorpb.HyperParams.proportion_clients', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='batch_size', full_name='aggregatorpb.HyperParams.batch_size', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='num_rounds', full_name='aggregatorpb.HyperParams.num_rounds', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='learning_rate', full_name='aggregatorpb.HyperParams.learning_rate', index=3,
      number=4, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='num_local_rounds', full_name='aggregatorpb.HyperParams.num_local_rounds', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=95,
  serialized_end=225,
)


_WORK = _descriptor.Descriptor(
  name='Work',
  full_name='aggregatorpb.Work',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='aggregatorpb.Work.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))),
    _descriptor.FieldDescriptor(
      name='num_examples', full_name='aggregatorpb.Work.num_examples', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='num_clients', full_name='aggregatorpb.Work.num_clients', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='epoch', full_name='aggregatorpb.Work.epoch', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='model', full_name='aggregatorpb.Work.model', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='hyper_params', full_name='aggregatorpb.Work.hyper_params', index=5,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=228,
  serialized_end=403,
)


_MODELID = _descriptor.Descriptor(
  name='ModelID',
  full_name='aggregatorpb.ModelID',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='domain', full_name='aggregatorpb.ModelID.domain', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='aggregatorpb.ModelID.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='id', full_name='aggregatorpb.ModelID.id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=405,
  serialized_end=456,
)


_GETWORKREQUEST = _descriptor.Descriptor(
  name='GetWorkRequest',
  full_name='aggregatorpb.GetWorkRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ids', full_name='aggregatorpb.GetWorkRequest.ids', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=458,
  serialized_end=516,
)


_REPORTWORKREQUEST = _descriptor.Descriptor(
  name='ReportWorkRequest',
  full_name='aggregatorpb.ReportWorkRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='work', full_name='aggregatorpb.ReportWorkRequest.work', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=518,
  serialized_end=577,
)


_REPORTWORKREPLY = _descriptor.Descriptor(
  name='ReportWorkReply',
  full_name='aggregatorpb.ReportWorkReply',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=579,
  serialized_end=596,
)


_PRODMODELREQUEST = _descriptor.Descriptor(
  name='ProdModelRequest',
  full_name='aggregatorpb.ProdModelRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ids', full_name='aggregatorpb.ProdModelRequest.ids', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=_descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=598,
  serialized_end=658,
)


_PRODMODELRESPONSE = _descriptor.Descriptor(
  name='ProdModelResponse',
  full_name='aggregatorpb.ProdModelResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='model_urls', full_name='aggregatorpb.ProdModelResponse.model_urls', index=0,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=660,
  serialized_end=699,
)

_WORK.fields_by_name['id'].message_type = _MODELID
_WORK.fields_by_name['hyper_params'].message_type = _HYPERPARAMS
_GETWORKREQUEST.fields_by_name['ids'].message_type = _MODELID
_REPORTWORKREQUEST.fields_by_name['work'].message_type = _WORK
_PRODMODELREQUEST.fields_by_name['ids'].message_type = _MODELID
DESCRIPTOR.message_types_by_name['HyperParams'] = _HYPERPARAMS
DESCRIPTOR.message_types_by_name['Work'] = _WORK
DESCRIPTOR.message_types_by_name['ModelID'] = _MODELID
DESCRIPTOR.message_types_by_name['GetWorkRequest'] = _GETWORKREQUEST
DESCRIPTOR.message_types_by_name['ReportWorkRequest'] = _REPORTWORKREQUEST
DESCRIPTOR.message_types_by_name['ReportWorkReply'] = _REPORTWORKREPLY
DESCRIPTOR.message_types_by_name['ProdModelRequest'] = _PRODMODELREQUEST
DESCRIPTOR.message_types_by_name['ProdModelResponse'] = _PRODMODELRESPONSE
DESCRIPTOR.enum_types_by_name['TrainingStatus'] = _TRAININGSTATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

HyperParams = _reflection.GeneratedProtocolMessageType('HyperParams', (_message.Message,), dict(
  DESCRIPTOR = _HYPERPARAMS,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.HyperParams)
  ))
_sym_db.RegisterMessage(HyperParams)

Work = _reflection.GeneratedProtocolMessageType('Work', (_message.Message,), dict(
  DESCRIPTOR = _WORK,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.Work)
  ))
_sym_db.RegisterMessage(Work)

ModelID = _reflection.GeneratedProtocolMessageType('ModelID', (_message.Message,), dict(
  DESCRIPTOR = _MODELID,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.ModelID)
  ))
_sym_db.RegisterMessage(ModelID)

GetWorkRequest = _reflection.GeneratedProtocolMessageType('GetWorkRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETWORKREQUEST,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.GetWorkRequest)
  ))
_sym_db.RegisterMessage(GetWorkRequest)

ReportWorkRequest = _reflection.GeneratedProtocolMessageType('ReportWorkRequest', (_message.Message,), dict(
  DESCRIPTOR = _REPORTWORKREQUEST,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.ReportWorkRequest)
  ))
_sym_db.RegisterMessage(ReportWorkRequest)

ReportWorkReply = _reflection.GeneratedProtocolMessageType('ReportWorkReply', (_message.Message,), dict(
  DESCRIPTOR = _REPORTWORKREPLY,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.ReportWorkReply)
  ))
_sym_db.RegisterMessage(ReportWorkReply)

ProdModelRequest = _reflection.GeneratedProtocolMessageType('ProdModelRequest', (_message.Message,), dict(
  DESCRIPTOR = _PRODMODELREQUEST,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.ProdModelRequest)
  ))
_sym_db.RegisterMessage(ProdModelRequest)

ProdModelResponse = _reflection.GeneratedProtocolMessageType('ProdModelResponse', (_message.Message,), dict(
  DESCRIPTOR = _PRODMODELRESPONSE,
  __module__ = 'aggregatorpb.aggregator_pb2'
  # @@protoc_insertion_point(class_scope:aggregatorpb.ProdModelResponse)
  ))
_sym_db.RegisterMessage(ProdModelResponse)


_WORK.fields_by_name['id'].has_options = True
_WORK.fields_by_name['id']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_WORK.fields_by_name['hyper_params'].has_options = True
_WORK.fields_by_name['hyper_params']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_GETWORKREQUEST.fields_by_name['ids'].has_options = True
_GETWORKREQUEST.fields_by_name['ids']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_REPORTWORKREQUEST.fields_by_name['work'].has_options = True
_REPORTWORKREQUEST.fields_by_name['work']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
_PRODMODELREQUEST.fields_by_name['ids'].has_options = True
_PRODMODELREQUEST.fields_by_name['ids']._options = _descriptor._ParseOptions(descriptor_pb2.FieldOptions(), _b('\310\336\037\000'))
try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class AggregatorStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.GetWork = channel.unary_stream(
          '/aggregatorpb.Aggregator/GetWork',
          request_serializer=GetWorkRequest.SerializeToString,
          response_deserializer=Work.FromString,
          )
      self.ReportWork = channel.unary_unary(
          '/aggregatorpb.Aggregator/ReportWork',
          request_serializer=ReportWorkRequest.SerializeToString,
          response_deserializer=ReportWorkReply.FromString,
          )
      self.ProdModel = channel.unary_unary(
          '/aggregatorpb.Aggregator/ProdModel',
          request_serializer=ProdModelRequest.SerializeToString,
          response_deserializer=ProdModelResponse.FromString,
          )


  class AggregatorServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def GetWork(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ReportWork(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ProdModel(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_AggregatorServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'GetWork': grpc.unary_stream_rpc_method_handler(
            servicer.GetWork,
            request_deserializer=GetWorkRequest.FromString,
            response_serializer=Work.SerializeToString,
        ),
        'ReportWork': grpc.unary_unary_rpc_method_handler(
            servicer.ReportWork,
            request_deserializer=ReportWorkRequest.FromString,
            response_serializer=ReportWorkReply.SerializeToString,
        ),
        'ProdModel': grpc.unary_unary_rpc_method_handler(
            servicer.ProdModel,
            request_deserializer=ProdModelRequest.FromString,
            response_serializer=ProdModelResponse.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'aggregatorpb.Aggregator', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaAggregatorServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def GetWork(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ReportWork(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ProdModel(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaAggregatorStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def GetWork(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    def ReportWork(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    ReportWork.future = None
    def ProdModel(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    ProdModel.future = None


  def beta_create_Aggregator_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('aggregatorpb.Aggregator', 'GetWork'): GetWorkRequest.FromString,
      ('aggregatorpb.Aggregator', 'ProdModel'): ProdModelRequest.FromString,
      ('aggregatorpb.Aggregator', 'ReportWork'): ReportWorkRequest.FromString,
    }
    response_serializers = {
      ('aggregatorpb.Aggregator', 'GetWork'): Work.SerializeToString,
      ('aggregatorpb.Aggregator', 'ProdModel'): ProdModelResponse.SerializeToString,
      ('aggregatorpb.Aggregator', 'ReportWork'): ReportWorkReply.SerializeToString,
    }
    method_implementations = {
      ('aggregatorpb.Aggregator', 'GetWork'): face_utilities.unary_stream_inline(servicer.GetWork),
      ('aggregatorpb.Aggregator', 'ProdModel'): face_utilities.unary_unary_inline(servicer.ProdModel),
      ('aggregatorpb.Aggregator', 'ReportWork'): face_utilities.unary_unary_inline(servicer.ReportWork),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Aggregator_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('aggregatorpb.Aggregator', 'GetWork'): GetWorkRequest.SerializeToString,
      ('aggregatorpb.Aggregator', 'ProdModel'): ProdModelRequest.SerializeToString,
      ('aggregatorpb.Aggregator', 'ReportWork'): ReportWorkRequest.SerializeToString,
    }
    response_deserializers = {
      ('aggregatorpb.Aggregator', 'GetWork'): Work.FromString,
      ('aggregatorpb.Aggregator', 'ProdModel'): ProdModelResponse.FromString,
      ('aggregatorpb.Aggregator', 'ReportWork'): ReportWorkReply.FromString,
    }
    cardinalities = {
      'GetWork': cardinality.Cardinality.UNARY_STREAM,
      'ProdModel': cardinality.Cardinality.UNARY_UNARY,
      'ReportWork': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'aggregatorpb.Aggregator', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)

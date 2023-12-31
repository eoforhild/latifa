// AUTO GENERATED FILE, DO NOT EDIT.
// Generated by `flutter_rust_bridge`@ 1.82.3.
// ignore_for_file: non_constant_identifier_names, unused_element, duplicate_ignore, directives_ordering, curly_braces_in_flow_control_structures, unnecessary_lambdas, slash_for_doc_comments, prefer_const_literals_to_create_immutables, implicit_dynamic_list_literal, duplicate_import, unused_import, unnecessary_import, prefer_single_quotes, prefer_const_constructors, use_super_parameters, always_use_package_imports, annotate_overrides, invalid_use_of_protected_member, constant_identifier_names, invalid_use_of_internal_member, prefer_is_empty, unnecessary_const

import "bridge_definitions.dart";
import 'dart:convert';
import 'dart:async';
import 'package:meta/meta.dart';
import 'package:flutter_rust_bridge/flutter_rust_bridge.dart';
import 'package:uuid/uuid.dart';

import 'dart:convert';
import 'dart:async';
import 'package:meta/meta.dart';
import 'package:flutter_rust_bridge/flutter_rust_bridge.dart';
import 'package:uuid/uuid.dart';

import 'dart:ffi' as ffi;

class NativeImpl implements Native {
  final NativePlatform _platform;
  factory NativeImpl(ExternalLibrary dylib) =>
      NativeImpl.raw(NativePlatform(dylib));

  /// Only valid on web/WASM platforms.
  factory NativeImpl.wasm(FutureOr<WasmModule> module) =>
      NativeImpl(module as ExternalLibrary);
  NativeImpl.raw(this._platform);
  Future<U8Array32> kdf({required Uint8List km, dynamic hint}) {
    var arg0 = _platform.api2wire_uint_8_list(km);
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_kdf(port_, arg0),
      parseSuccessData: _wire2api_u8_array_32,
      parseErrorData: null,
      constMeta: kKdfConstMeta,
      argValues: [km],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kKdfConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "kdf",
        argNames: ["km"],
      );

  Future<void> generateKeysAndDump({dynamic hint}) {
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_generate_keys_and_dump(port_),
      parseSuccessData: _wire2api_unit,
      parseErrorData: null,
      constMeta: kGenerateKeysAndDumpConstMeta,
      argValues: [],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kGenerateKeysAndDumpConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "generate_keys_and_dump",
        argNames: [],
      );

  Future<bool> registerAndPublish({required String regForm, dynamic hint}) {
    var arg0 = _platform.api2wire_String(regForm);
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) =>
          _platform.inner.wire_register_and_publish(port_, arg0),
      parseSuccessData: _wire2api_bool,
      parseErrorData: null,
      constMeta: kRegisterAndPublishConstMeta,
      argValues: [regForm],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kRegisterAndPublishConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "register_and_publish",
        argNames: ["regForm"],
      );

  Future<bool> login({required String logForm, dynamic hint}) {
    var arg0 = _platform.api2wire_String(logForm);
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_login(port_, arg0),
      parseSuccessData: _wire2api_bool,
      parseErrorData: null,
      constMeta: kLoginConstMeta,
      argValues: [logForm],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kLoginConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "login",
        argNames: ["logForm"],
      );

  Future<void> logout({dynamic hint}) {
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_logout(port_),
      parseSuccessData: _wire2api_unit,
      parseErrorData: null,
      constMeta: kLogoutConstMeta,
      argValues: [],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kLogoutConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "logout",
        argNames: [],
      );

  Future<bool> requestConnection({required String email, dynamic hint}) {
    var arg0 = _platform.api2wire_String(email);
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_request_connection(port_, arg0),
      parseSuccessData: _wire2api_bool,
      parseErrorData: null,
      constMeta: kRequestConnectionConstMeta,
      argValues: [email],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kRequestConnectionConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "request_connection",
        argNames: ["email"],
      );

  Future<bool> pendingRequests({dynamic hint}) {
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_pending_requests(port_),
      parseSuccessData: _wire2api_bool,
      parseErrorData: null,
      constMeta: kPendingRequestsConstMeta,
      argValues: [],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kPendingRequestsConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "pending_requests",
        argNames: [],
      );

  Future<void> approvedRequests({dynamic hint}) {
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) => _platform.inner.wire_approved_requests(port_),
      parseSuccessData: _wire2api_unit,
      parseErrorData: null,
      constMeta: kApprovedRequestsConstMeta,
      argValues: [],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kApprovedRequestsConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "approved_requests",
        argNames: [],
      );

  Future<void> fetchKeysHandshake({required String reqId, dynamic hint}) {
    var arg0 = _platform.api2wire_String(reqId);
    return _platform.executeNormal(FlutterRustBridgeTask(
      callFfi: (port_) =>
          _platform.inner.wire_fetch_keys_handshake(port_, arg0),
      parseSuccessData: _wire2api_unit,
      parseErrorData: null,
      constMeta: kFetchKeysHandshakeConstMeta,
      argValues: [reqId],
      hint: hint,
    ));
  }

  FlutterRustBridgeTaskConstMeta get kFetchKeysHandshakeConstMeta =>
      const FlutterRustBridgeTaskConstMeta(
        debugName: "fetch_keys_handshake",
        argNames: ["reqId"],
      );

  void dispose() {
    _platform.dispose();
  }
// Section: wire2api

  bool _wire2api_bool(dynamic raw) {
    return raw as bool;
  }

  int _wire2api_u8(dynamic raw) {
    return raw as int;
  }

  U8Array32 _wire2api_u8_array_32(dynamic raw) {
    return U8Array32(_wire2api_uint_8_list(raw));
  }

  Uint8List _wire2api_uint_8_list(dynamic raw) {
    return raw as Uint8List;
  }

  void _wire2api_unit(dynamic raw) {
    return;
  }
}

// Section: api2wire

@protected
int api2wire_u8(int raw) {
  return raw;
}

// Section: finalizer

class NativePlatform extends FlutterRustBridgeBase<NativeWire> {
  NativePlatform(ffi.DynamicLibrary dylib) : super(NativeWire(dylib));

// Section: api2wire

  @protected
  ffi.Pointer<wire_uint_8_list> api2wire_String(String raw) {
    return api2wire_uint_8_list(utf8.encoder.convert(raw));
  }

  @protected
  ffi.Pointer<wire_uint_8_list> api2wire_uint_8_list(Uint8List raw) {
    final ans = inner.new_uint_8_list_0(raw.length);
    ans.ref.ptr.asTypedList(raw.length).setAll(0, raw);
    return ans;
  }
// Section: finalizer

// Section: api_fill_to_wire
}

// ignore_for_file: camel_case_types, non_constant_identifier_names, avoid_positional_boolean_parameters, annotate_overrides, constant_identifier_names

// AUTO GENERATED FILE, DO NOT EDIT.
//
// Generated by `package:ffigen`.
// ignore_for_file: type=lint

/// generated by flutter_rust_bridge
class NativeWire implements FlutterRustBridgeWireBase {
  @internal
  late final dartApi = DartApiDl(init_frb_dart_api_dl);

  /// Holds the symbol lookup function.
  final ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
      _lookup;

  /// The symbols are looked up in [dynamicLibrary].
  NativeWire(ffi.DynamicLibrary dynamicLibrary)
      : _lookup = dynamicLibrary.lookup;

  /// The symbols are looked up with [lookup].
  NativeWire.fromLookup(
      ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
          lookup)
      : _lookup = lookup;

  void store_dart_post_cobject(
    DartPostCObjectFnType ptr,
  ) {
    return _store_dart_post_cobject(
      ptr,
    );
  }

  late final _store_dart_post_cobjectPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(DartPostCObjectFnType)>>(
          'store_dart_post_cobject');
  late final _store_dart_post_cobject = _store_dart_post_cobjectPtr
      .asFunction<void Function(DartPostCObjectFnType)>();

  Object get_dart_object(
    int ptr,
  ) {
    return _get_dart_object(
      ptr,
    );
  }

  late final _get_dart_objectPtr =
      _lookup<ffi.NativeFunction<ffi.Handle Function(ffi.UintPtr)>>(
          'get_dart_object');
  late final _get_dart_object =
      _get_dart_objectPtr.asFunction<Object Function(int)>();

  void drop_dart_object(
    int ptr,
  ) {
    return _drop_dart_object(
      ptr,
    );
  }

  late final _drop_dart_objectPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.UintPtr)>>(
          'drop_dart_object');
  late final _drop_dart_object =
      _drop_dart_objectPtr.asFunction<void Function(int)>();

  int new_dart_opaque(
    Object handle,
  ) {
    return _new_dart_opaque(
      handle,
    );
  }

  late final _new_dart_opaquePtr =
      _lookup<ffi.NativeFunction<ffi.UintPtr Function(ffi.Handle)>>(
          'new_dart_opaque');
  late final _new_dart_opaque =
      _new_dart_opaquePtr.asFunction<int Function(Object)>();

  int init_frb_dart_api_dl(
    ffi.Pointer<ffi.Void> obj,
  ) {
    return _init_frb_dart_api_dl(
      obj,
    );
  }

  late final _init_frb_dart_api_dlPtr =
      _lookup<ffi.NativeFunction<ffi.IntPtr Function(ffi.Pointer<ffi.Void>)>>(
          'init_frb_dart_api_dl');
  late final _init_frb_dart_api_dl = _init_frb_dart_api_dlPtr
      .asFunction<int Function(ffi.Pointer<ffi.Void>)>();

  void wire_kdf(
    int port_,
    ffi.Pointer<wire_uint_8_list> km,
  ) {
    return _wire_kdf(
      port_,
      km,
    );
  }

  late final _wire_kdfPtr = _lookup<
      ffi.NativeFunction<
          ffi.Void Function(
              ffi.Int64, ffi.Pointer<wire_uint_8_list>)>>('wire_kdf');
  late final _wire_kdf = _wire_kdfPtr
      .asFunction<void Function(int, ffi.Pointer<wire_uint_8_list>)>();

  void wire_generate_keys_and_dump(
    int port_,
  ) {
    return _wire_generate_keys_and_dump(
      port_,
    );
  }

  late final _wire_generate_keys_and_dumpPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.Int64)>>(
          'wire_generate_keys_and_dump');
  late final _wire_generate_keys_and_dump =
      _wire_generate_keys_and_dumpPtr.asFunction<void Function(int)>();

  void wire_register_and_publish(
    int port_,
    ffi.Pointer<wire_uint_8_list> reg_form,
  ) {
    return _wire_register_and_publish(
      port_,
      reg_form,
    );
  }

  late final _wire_register_and_publishPtr = _lookup<
      ffi.NativeFunction<
          ffi.Void Function(ffi.Int64,
              ffi.Pointer<wire_uint_8_list>)>>('wire_register_and_publish');
  late final _wire_register_and_publish = _wire_register_and_publishPtr
      .asFunction<void Function(int, ffi.Pointer<wire_uint_8_list>)>();

  void wire_login(
    int port_,
    ffi.Pointer<wire_uint_8_list> log_form,
  ) {
    return _wire_login(
      port_,
      log_form,
    );
  }

  late final _wire_loginPtr = _lookup<
      ffi.NativeFunction<
          ffi.Void Function(
              ffi.Int64, ffi.Pointer<wire_uint_8_list>)>>('wire_login');
  late final _wire_login = _wire_loginPtr
      .asFunction<void Function(int, ffi.Pointer<wire_uint_8_list>)>();

  void wire_logout(
    int port_,
  ) {
    return _wire_logout(
      port_,
    );
  }

  late final _wire_logoutPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.Int64)>>('wire_logout');
  late final _wire_logout = _wire_logoutPtr.asFunction<void Function(int)>();

  void wire_request_connection(
    int port_,
    ffi.Pointer<wire_uint_8_list> email,
  ) {
    return _wire_request_connection(
      port_,
      email,
    );
  }

  late final _wire_request_connectionPtr = _lookup<
      ffi.NativeFunction<
          ffi.Void Function(ffi.Int64,
              ffi.Pointer<wire_uint_8_list>)>>('wire_request_connection');
  late final _wire_request_connection = _wire_request_connectionPtr
      .asFunction<void Function(int, ffi.Pointer<wire_uint_8_list>)>();

  void wire_pending_requests(
    int port_,
  ) {
    return _wire_pending_requests(
      port_,
    );
  }

  late final _wire_pending_requestsPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.Int64)>>(
          'wire_pending_requests');
  late final _wire_pending_requests =
      _wire_pending_requestsPtr.asFunction<void Function(int)>();

  void wire_approved_requests(
    int port_,
  ) {
    return _wire_approved_requests(
      port_,
    );
  }

  late final _wire_approved_requestsPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.Int64)>>(
          'wire_approved_requests');
  late final _wire_approved_requests =
      _wire_approved_requestsPtr.asFunction<void Function(int)>();

  void wire_fetch_keys_handshake(
    int port_,
    ffi.Pointer<wire_uint_8_list> req_id,
  ) {
    return _wire_fetch_keys_handshake(
      port_,
      req_id,
    );
  }

  late final _wire_fetch_keys_handshakePtr = _lookup<
      ffi.NativeFunction<
          ffi.Void Function(ffi.Int64,
              ffi.Pointer<wire_uint_8_list>)>>('wire_fetch_keys_handshake');
  late final _wire_fetch_keys_handshake = _wire_fetch_keys_handshakePtr
      .asFunction<void Function(int, ffi.Pointer<wire_uint_8_list>)>();

  ffi.Pointer<wire_uint_8_list> new_uint_8_list_0(
    int len,
  ) {
    return _new_uint_8_list_0(
      len,
    );
  }

  late final _new_uint_8_list_0Ptr = _lookup<
          ffi
          .NativeFunction<ffi.Pointer<wire_uint_8_list> Function(ffi.Int32)>>(
      'new_uint_8_list_0');
  late final _new_uint_8_list_0 = _new_uint_8_list_0Ptr
      .asFunction<ffi.Pointer<wire_uint_8_list> Function(int)>();

  void free_WireSyncReturn(
    WireSyncReturn ptr,
  ) {
    return _free_WireSyncReturn(
      ptr,
    );
  }

  late final _free_WireSyncReturnPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(WireSyncReturn)>>(
          'free_WireSyncReturn');
  late final _free_WireSyncReturn =
      _free_WireSyncReturnPtr.asFunction<void Function(WireSyncReturn)>();
}

final class _Dart_Handle extends ffi.Opaque {}

final class wire_uint_8_list extends ffi.Struct {
  external ffi.Pointer<ffi.Uint8> ptr;

  @ffi.Int32()
  external int len;
}

typedef DartPostCObjectFnType = ffi.Pointer<
    ffi.NativeFunction<
        ffi.Bool Function(DartPort port_id, ffi.Pointer<ffi.Void> message)>>;
typedef DartPort = ffi.Int64;

// GENERATED CODE - DO NOT MODIFY BY HAND
// This code was generated by ObjectBox. To update it run the generator again
// with `dart run build_runner build`.
// See also https://docs.objectbox.io/getting-started#generate-objectbox-code

// ignore_for_file: camel_case_types, depend_on_referenced_packages
// coverage:ignore-file

import 'dart:typed_data';

import 'package:flat_buffers/flat_buffers.dart' as fb;
import 'package:objectbox/internal.dart'
    as obx_int; // generated code can access "internal" functionality
import 'package:objectbox/objectbox.dart' as obx;
import 'package:objectbox_flutter_libs/objectbox_flutter_libs.dart';

import '../../../../core/network/data/user_credentials.dart';

export 'package:objectbox/objectbox.dart'; // so that callers only have to import this file

final _entities = <obx_int.ModelEntity>[
  obx_int.ModelEntity(
    id: const obx_int.IdUid(1, 2265316646599911411),
    name: 'UserCredentials',
    lastPropertyId: const obx_int.IdUid(2, 346814284049679867),
    flags: 0,
    properties: <obx_int.ModelProperty>[
      obx_int.ModelProperty(
        id: const obx_int.IdUid(1, 5840476384718239313),
        name: 'id',
        type: 6,
        flags: 129,
      ),
      obx_int.ModelProperty(
        id: const obx_int.IdUid(2, 346814284049679867),
        name: 'token',
        type: 9,
        flags: 0,
      ),
    ],
    relations: <obx_int.ModelRelation>[],
    backlinks: <obx_int.ModelBacklink>[],
  ),
];

/// Shortcut for [obx.Store.new] that passes [getObjectBoxModel] and for Flutter
/// apps by default a [directory] using `defaultStoreDirectory()` from the
/// ObjectBox Flutter library.
///
/// Note: for desktop apps it is recommended to specify a unique [directory].
///
/// See [obx.Store.new] for an explanation of all parameters.
///
/// For Flutter apps, also calls `loadObjectBoxLibraryAndroidCompat()` from
/// the ObjectBox Flutter library to fix loading the native ObjectBox library
/// on Android 6 and older.
Future<obx.Store> openStore({
  String? directory,
  int? maxDBSizeInKB,
  int? maxDataSizeInKB,
  int? fileMode,
  int? maxReaders,
  bool queriesCaseSensitiveDefault = true,
  String? macosApplicationGroup,
}) async {
  await loadObjectBoxLibraryAndroidCompat();
  return obx.Store(
    getObjectBoxModel(),
    directory: directory ?? (await defaultStoreDirectory()).path,
    maxDBSizeInKB: maxDBSizeInKB,
    maxDataSizeInKB: maxDataSizeInKB,
    fileMode: fileMode,
    maxReaders: maxReaders,
    queriesCaseSensitiveDefault: queriesCaseSensitiveDefault,
    macosApplicationGroup: macosApplicationGroup,
  );
}

/// Returns the ObjectBox model definition for this project for use with
/// [obx.Store.new].
obx_int.ModelDefinition getObjectBoxModel() {
  final model = obx_int.ModelInfo(
    entities: _entities,
    lastEntityId: const obx_int.IdUid(1, 2265316646599911411),
    lastIndexId: const obx_int.IdUid(0, 0),
    lastRelationId: const obx_int.IdUid(0, 0),
    lastSequenceId: const obx_int.IdUid(0, 0),
    retiredEntityUids: const [],
    retiredIndexUids: const [],
    retiredPropertyUids: const [],
    retiredRelationUids: const [],
    modelVersion: 5,
    modelVersionParserMinimum: 5,
    version: 1,
  );

  final bindings = <Type, obx_int.EntityDefinition>{
    UserCredentials: obx_int.EntityDefinition<UserCredentials>(
      model: _entities[0],
      toOneRelations: (UserCredentials object) => [],
      toManyRelations: (UserCredentials object) => {},
      getId: (UserCredentials object) => object.id,
      setId: (UserCredentials object, int id) {
        object.id = id;
      },
      objectToFB: (UserCredentials object, fb.Builder fbb) {
        final tokenOffset = fbb.writeString(object.token);
        fbb.startTable(3);
        fbb.addInt64(0, object.id);
        fbb.addOffset(1, tokenOffset);
        fbb.finish(fbb.endTable());
        return object.id;
      },
      objectFromFB: (obx.Store store, ByteData fbData) {
        final buffer = fb.BufferContext(fbData);
        final rootOffset = buffer.derefObject(0);

        final object = UserCredentials()
          ..id = const fb.Int64Reader().vTableGet(buffer, rootOffset, 4, 0)
          ..token = const fb.StringReader(
            asciiOptimization: true,
          ).vTableGet(buffer, rootOffset, 6, '');

        return object;
      },
    ),
  };

  return obx_int.ModelDefinition(model, bindings);
}

/// [UserCredentials] entity fields to define ObjectBox queries.
class UserCredentials_ {
  /// See [UserCredentials.id].
  static final id = obx.QueryIntegerProperty<UserCredentials>(
    _entities[0].properties[0],
  );

  /// See [UserCredentials.token].
  static final token = obx.QueryStringProperty<UserCredentials>(
    _entities[0].properties[1],
  );
}

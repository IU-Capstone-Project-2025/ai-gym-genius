import 'dart:convert';
import 'dart:math';
import 'package:flutter/foundation.dart';

import '../../domain/entities/exercise_info_entity.dart';

extension MuscleGroupExtension on MuscleGroup {
  // static method to convert string to enum
  static MuscleGroup fromMap(String value) {
    return MuscleGroup.values.firstWhere(
      (e) => e.name.toLowerCase() == value.toLowerCase(),
      orElse: () => throw Exception("Unknown muscle group: $value"),
    );
  }

  // Returns random muscleGroup
  static MuscleGroup random() {
    return MuscleGroup.values[Random().nextInt(MuscleGroup.values.length)];
  }
}

class ExerciseInfoDTO {
  final int id;
  final String name;
  final String? description;
  final String? imagePath;
  final String? url;
  List<String>? muscleGroups;

  ExerciseInfoDTO({
    required this.id,
    required this.name,
    this.description,
    this.imagePath,
    this.url,
    this.muscleGroups,
  });

  ExerciseInfoDTO copyWith({
    int? id,
    String? name,
    String? description,
    String? imagePath,
    List<String>? muscleGroups,
  }) {
    return ExerciseInfoDTO(
      id: id ?? this.id,
      name: name ?? this.name,
      description: description ?? this.description,
      imagePath: imagePath ?? this.imagePath,
      muscleGroups: muscleGroups ?? this.muscleGroups,
    );
  }

  factory ExerciseInfoDTO.fromEntity(ExerciseInfoEntity entity) {
    return ExerciseInfoDTO(
      id: entity.id,
      name: entity.name,
      description: entity.description,
      imagePath: entity.imagePath,
      url: entity.url,
      muscleGroups: entity.muscleGroups?.map((mg) => mg.name).toList(),
    );
  }

  ExerciseInfoEntity toEntity() {
    return ExerciseInfoEntity(
      id: id,
      name: name,
      description: description,
      imagePath: imagePath,
      url: url,
      muscleGroups: muscleGroups
          ?.map((str) => MuscleGroupExtension.fromMap(str))
          .toList(),
    );
  }

  factory ExerciseInfoDTO.fromSQLMap(Map<String, dynamic> map) {
    return ExerciseInfoDTO(
      id: map['id'] as int,
      name: map['name'] as String,
      description: map['description'] as String,
      imagePath: map['image_path'] as String,
      muscleGroups: (map['muscle_groups'] as List<dynamic>?)
          ?.map((mg) => mg as String)
          .toList(),
    );
  }

  factory ExerciseInfoDTO.fromJsonMap(Map<String, dynamic> map) {
    return ExerciseInfoDTO(
      id: map['id'] as int,
      name: map['name'] as String,
      description: map['description'] as String,
      imagePath: map['imagePath'] as String,
      muscleGroups: (map['muscleGroups'] as List<dynamic>?)
          ?.map((mg) => mg as String)
          .toList(),
    );
  }

  @override
  String toString() {
    return 'ExerciseInfoDto(id: $id, name: $name, description: $description, imagePath: $imagePath, muscleGroups: $muscleGroups)';
  }

  @override
  bool operator ==(covariant ExerciseInfoDTO other) {
    if (identical(this, other)) return true;

    return other.id == id &&
        other.name == name &&
        other.description == description &&
        other.imagePath == imagePath &&
        listEquals(other.muscleGroups, muscleGroups);
  }

  @override
  int get hashCode {
    return id.hashCode ^
        name.hashCode ^
        description.hashCode ^
        imagePath.hashCode ^
        muscleGroups.hashCode;
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'name': name,
      'description': description,
      'image_path': imagePath,
      'muscle_groups': muscleGroups,
    };
  }

  Map<String, dynamic> toCamelMap() {
    return <String, dynamic>{
      'id': id,
      'name': name,
      'description': description,
      'imagePath': imagePath,
      'muscleGroups': muscleGroups,
    };
  }

  String toJson() => json.encode(toMap());

  factory ExerciseInfoDTO.fromJson(String source) =>
      ExerciseInfoDTO.fromSQLMap(json.decode(source) as Map<String, dynamic>);
}

// ignore_for_file: public_member_api_docs, sort_constructors_first
import 'package:flutter/foundation.dart';

// That is the class that is going to be loaded from
// A dataset of exercises.

enum MuscleGroup {
  biceps,
  triceps,
  back,
  chest,
  legs,
  shoulders,
}

class ExerciseInfo {
  final int id;
  final String name;
  final String description;
  final String imagePath;
  List<MuscleGroup>? muscleGroups = [];

  ExerciseInfo({
    required this.id,
    required this.name,
    required this.description,
    required this.imagePath,
    this.muscleGroups,
  });

  ExerciseInfo copyWith({
    int? id,
    String? name,
    String? description,
    String? imagePath,
    List<MuscleGroup>? muscleGroups,
  }) {
    return ExerciseInfo(
      id: id ?? this.id,
      name: name ?? this.name,
      description: description ?? this.description,
      imagePath: imagePath ?? this.imagePath,
      muscleGroups: muscleGroups ?? this.muscleGroups,
    );
  }

  factory ExerciseInfo.fromMap(Map<String, dynamic> map) {
    List<String> muscleGroupsFromMap = map["muscleGroup"].cast<String>();
    List<MuscleGroup> realMuscleGroups = [];

    for (String elem in muscleGroupsFromMap) {
      switch (elem.toLowerCase()) {
        case "chest":
          realMuscleGroups.add(MuscleGroup.chest);
          break;
        case "back":
          realMuscleGroups.add(MuscleGroup.back);
          break;
        case "legs":
          realMuscleGroups.add(MuscleGroup.legs);
          break;
        case "shoulders":
          realMuscleGroups.add(MuscleGroup.shoulders);
          break;
        case "biceps":
          realMuscleGroups.add(MuscleGroup.biceps);
          break;
        case "triceps":
          realMuscleGroups.add(MuscleGroup.triceps);
          break;
        default:
          throw Exception("Unknown Json Format");
      }
    }

    return ExerciseInfo(
      id: map['id'] as int,
      name: map['name'] as String,
      description: map['description'] as String,
      imagePath: map['imagePath'] as String,
      muscleGroups: realMuscleGroups,
    );
  }

  @override
  String toString() {
    return 'ExerciseInfo(id: $id, name: $name, description: $description, imagePath: $imagePath, muscleGroups: $muscleGroups)';
  }

  @override
  bool operator ==(covariant ExerciseInfo other) {
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
}
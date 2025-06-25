enum MuscleGroup {
  biceps,
  triceps,
  back,
  chest,
  legs,
  shoulders,
}

class ExerciseInfoEntity {
  final int id;
  String name;
  String? description;
  String? imagePath;
  String? url;
  List<MuscleGroup>? muscleGroups;
  ExerciseInfoEntity({
    required this.id,
    required this.name,
    this.description,
    this.imagePath,
    this.url,
    this.muscleGroups,
  });
}

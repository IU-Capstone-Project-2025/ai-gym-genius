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
  final String name;
  final String? description;
  final String? imagePath;
  final String? url;
  final List<MuscleGroup>? muscleGroups;
  const ExerciseInfoEntity({
    required this.id,
    required this.name,
    this.description,
    this.imagePath,
    this.url,
    this.muscleGroups,
  });
}

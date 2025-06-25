import 'exercise_entity.dart';

class WorkoutEntity {
  final int id;
  final Duration duration;
  final DateTime startTime;
  final List<ExerciseEntity> exercises;
  final String? description;
  final double? weight;

  WorkoutEntity({
    required this.id,
    required this.duration,
    required this.startTime,
    required this.exercises,
    this.description,
    this.weight,
  });
}

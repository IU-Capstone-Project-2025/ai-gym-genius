import 'package:equatable/equatable.dart';

import '/core/domain/entities/exercise_entity.dart';
import '/core/domain/entities/workout_entity.dart';

enum AddExerciseStatus { initial, success, duplicate }

final class TrainingState extends Equatable {
  const TrainingState({
    this.exercises = const [],
    this.addExerciseStatus = AddExerciseStatus.initial,
    this.workout,
    this.review,
  });

  final List<ExerciseEntity> exercises;
  final AddExerciseStatus addExerciseStatus;
  final WorkoutEntity? workout;
  final String? review;

  @override
  List<Object?> get props => [
        exercises,
        addExerciseStatus,
        workout,
        review,
      ];

  TrainingState copyWith({
    List<ExerciseEntity>? exercises,
    AddExerciseStatus? addStatus,
    WorkoutEntity? workout,
    String? review,
  }) {
    return TrainingState(
      exercises: exercises ?? this.exercises,
      addExerciseStatus: addStatus ?? addExerciseStatus,
      workout: workout ?? workout,
      review: review ?? review,
    );
  }
}

import '/core/domain/entities/exercise_info_entity.dart';

sealed class TrainingEvent {
  const TrainingEvent();
}

class AddExercise extends TrainingEvent {
  final ExerciseInfoEntity info;
  const AddExercise(this.info);
}

class RemoveExercise extends TrainingEvent {
  final int exerciseID;
  const RemoveExercise(this.exerciseID);
}

class SubmitTraining extends TrainingEvent {
  final Duration workoutDuration;
  const SubmitTraining(this.workoutDuration);
}

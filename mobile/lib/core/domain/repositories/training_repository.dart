import '/core/domain/entities/workout_entity.dart';

abstract interface class TrainingRepository {
  Future<void> saveWorkout(WorkoutEntity entity);
  Future<List<WorkoutEntity>> fetchWorkouts();
  Future<WorkoutEntity?> fetchWorkout(int workoutId);
  Future<String> getAIReview(WorkoutEntity workout);
}
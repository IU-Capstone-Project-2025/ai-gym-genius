import '/core/domain/entities/workout_entity.dart';
import '/core/domain/repositories/training_repository.dart';

class TrainingRepositoryImpl implements TrainingRepository {
  @override
  Future<WorkoutEntity?> fetchWorkout(int workoutId) {
    // TODO: implement fetchWorkout
    throw UnimplementedError();
  }

  @override
  Future<List<WorkoutEntity>> fetchWorkouts() {
    // TODO: implement fetchWorkouts
    throw UnimplementedError();
  }

  @override
  Future<String> getAIReview(WorkoutEntity workout) {
    // TODO: implement getAIReview
    throw UnimplementedError();
  }

  @override
  Future<void> saveWorkout(WorkoutEntity entity) {
    // TODO: implement saveWorkout
    throw UnimplementedError();
  }
}

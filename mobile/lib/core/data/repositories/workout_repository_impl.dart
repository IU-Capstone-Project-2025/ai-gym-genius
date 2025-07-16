import '/core/domain/entities/workout_entity.dart';
import '../../domain/repositories/workout_repository.dart';

class WorkoutRepositoryImpl implements WorkoutRepository {
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

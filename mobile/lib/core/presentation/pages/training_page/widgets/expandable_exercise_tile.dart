import 'dart:ui' show ImageFilter;
import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/services.dart' show FilteringTextInputFormatter;

import '/core/domain/entities/exercise_entity.dart';
import '/core/domain/entities/exercise_set_entity.dart';

class ExpandableExerciseTile extends StatefulWidget {
  final ExerciseEntity exercise;
  final void Function(ExerciseEntity) onDeleteCallback;

  const ExpandableExerciseTile({
    super.key,
    required this.exercise,
    required this.onDeleteCallback,
  });

  @override
  State<ExpandableExerciseTile> createState() => _ExpandableExerciseTileState();
}

class _ExpandableExerciseTileState extends State<ExpandableExerciseTile> {
  final Duration _duration = Durations.short3;
  static const double _padding = 16.0 + 8;
  bool isExpanded = false;

  void addSetToExercise(ExerciseSetEntity exerciseSet) {
    setState(() {
      widget.exercise.sets.add(exerciseSet);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.symmetric(horizontal: _padding),
      child: Column(
        children: [
          _buildMainContent(),
          _buildExpandablePart(),
        ],
      ),
    );
  }

  /// Main content of tile, that is always visible
  /// TODO: make more adaptive
  Widget _buildMainContent() {
    return GestureDetector(
      onTap: () {
        setState(() {
          isExpanded = !isExpanded;
        });
      },
      child: AnimatedContainer(
        curve: Curves.easeIn,
        duration: _duration,
        height: 80,
        margin: isExpanded
            ? EdgeInsets.only(top: 16, bottom: 8)
            : EdgeInsets.symmetric(vertical: 16),
        child: ClipRRect(
          borderRadius: BorderRadius.all(Radius.circular(18)),
          child: Stack(
            children: [
              if (widget.exercise.exerciseInfo.imagePath != "")
                Image.asset(
                  widget.exercise.exerciseInfo.imagePath!,
                  width: double.infinity,
                  height: double.infinity,
                  fit: BoxFit.cover,
                )
              else
                Container(
                  width: double.infinity,
                  height: double.infinity,
                  color: Colors.grey.shade300,
                  child: const Icon(
                    CupertinoIcons.photo,
                    size: 64,
                    color: Colors.white54,
                  ),
                ),
              // Blur effect
              BackdropFilter(
                filter: ImageFilter.blur(sigmaX: 3, sigmaY: 3),
                child: Container(
                  color: Theme.of(context)
                      .colorScheme
                      .secondaryContainer
                      .withAlpha(150),
                ),
              ),
              // Content
              _buildMainContentLayout(),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildMainContentLayout() => Padding(
        padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            // Text with Exercise name
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  maxLines: 1,
                  overflow: TextOverflow.ellipsis,
                  widget.exercise.exerciseInfo.name,
                  style: TextStyle().copyWith(
                    fontSize: 24,
                  ),
                ),
                Text('${widget.exercise.sets.length} Set(s)')
              ],
            ),
            // Row with small icons
            Row(
              mainAxisAlignment: MainAxisAlignment.end,
              children: [
                CupertinoButton(
                  padding: EdgeInsets.all(0),
                  onPressed: () => widget.onDeleteCallback(widget.exercise),
                  child: Icon(CupertinoIcons.trash),
                ),
                // Animated chevron
                isExpanded
                    ? Icon(CupertinoIcons.chevron_down)
                    : Icon(CupertinoIcons.chevron_forward),
              ],
            ),
          ],
        ),
      );

  /// Part which expands on tap
  Widget _buildExpandablePart() {
    return Column(
      children: [
        SetAdder(
          duration: _duration,
          isExpanded: isExpanded,
          setAdderCallback: addSetToExercise,
        ),
        _buildAmountText(),
        ...List.generate(
          widget.exercise.sets.length,
          (value) {
            ExerciseSetEntity exerciseSet = widget.exercise.sets[value];
            return _buildExerciseSetCard(exerciseSet, value);
          },
        ),
        // This widget is just cooking.
        AnimatedContainer(
          duration: _duration,
          height: isExpanded ? 8 : 0,
        )
      ],
    );
  }

  /// Text showing if sets added in Expandable part
  Widget _buildAmountText() {
    return AnimatedContainer(
      duration: _duration,
      alignment: Alignment.bottomLeft,
      height: isExpanded ? 32 : 0,
      child: AnimatedCrossFade(
          duration: _duration,
          firstChild: SizedBox.shrink(),
          secondChild: Padding(
            padding: EdgeInsets.only(top: 0),
            child: Text(
              widget.exercise.sets.isEmpty
                  ? "No Sets Completed"
                  : 'Sets Completed',
              style: CupertinoTheme.of(context)
                  .textTheme
                  .navTitleTextStyle
                  .copyWith(color: Theme.of(context).colorScheme.onSecondary),
            ),
          ),
          crossFadeState: isExpanded
              ? CrossFadeState.showSecond
              : CrossFadeState.showFirst),
    );
  }

  /// Small Orange Card that shows set information
  Widget _buildExerciseSetCard(ExerciseSetEntity exerciseSet, int index) {
    return AnimatedContainer(
      curve: Curves.decelerate,
      margin: isExpanded ? EdgeInsets.symmetric(vertical: 8) : EdgeInsets.zero,
      decoration: BoxDecoration(
        color: Theme.of(context).colorScheme.secondaryContainer,
        borderRadius: BorderRadius.circular(12),
      ),
      duration: _duration,
      height: isExpanded ? 40 : 0,
      child: AnimatedCrossFade(
        crossFadeState:
            isExpanded ? CrossFadeState.showSecond : CrossFadeState.showFirst,
        duration: _duration,
        firstChild: SizedBox.shrink(),
        secondChild: Center(
          child: DefaultTextStyle(
            style: TextStyle().copyWith(color: Colors.white),
            child: Row(
              children: [
                SizedBox(
                  width: 16,
                ),
                Text(
                  'Set ${index + 1}',
                  style: TextStyle().copyWith(fontWeight: FontWeight.w600),
                ),
                Expanded(
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Text('${exerciseSet.reps} time(s)'),
                      Text('${exerciseSet.weight.toString()} kgs'),
                    ],
                  ),
                ),
                Row(
                  children: [
                    CupertinoButton(
                      // color:
                      //     Theme.of(context).colorScheme.primary.withAlpha(100),
                      padding: EdgeInsets.zero,
                      onPressed: () {
                        setState(() {
                          widget.exercise.sets.removeAt(index);
                        });
                      },
                      child: Icon(
                        CupertinoIcons.delete,
                      ),
                    ),
                    SizedBox(width: 16),
                  ],
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}

/// Two Fields + Button that adds set
class SetAdder extends StatefulWidget {
  const SetAdder({
    super.key,
    required Duration duration,
    required this.isExpanded,
    required this.setAdderCallback,
  }) : _duration = duration;

  final Duration _duration;
  final bool isExpanded;
  final Function(ExerciseSetEntity) setAdderCallback;

  @override
  State<SetAdder> createState() => _SetAdderState();
}

class _SetAdderState extends State<SetAdder> {
  late TextEditingController repsController;
  late TextEditingController weightController;

  bool repsControllerHasError = false;
  bool weightControllerHasError = false;

  @override
  void initState() {
    super.initState();
    repsController = TextEditingController();
    weightController = TextEditingController();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedContainer(
      duration: widget._duration,
      height: widget.isExpanded ? 40 : 0,
      child: widget.isExpanded
          ? Row(
              children: [
                Expanded(
                  flex: 3,
                  child: Row(
                    children: [
                      Expanded(
                        child: CupertinoTextField(
                          onChanged: (value) {
                            setState(() {
                              repsControllerHasError = false;
                            });
                          },
                          decoration: getTextFieldDecoration(
                            repsControllerHasError,
                          ),
                          inputFormatters: [
                            FilteringTextInputFormatter.allow(RegExp(r'[0-9]'))
                          ],
                          controller: repsController,
                          keyboardType: TextInputType.number,
                          placeholder: "Reps",
                          style: TextStyle().copyWith(
                            color:
                                Theme.of(context).colorScheme.primaryContainer,
                          ),
                        ),
                      ),
                      SizedBox(width: 8),
                      Expanded(
                        child: CupertinoTextField(
                          onChanged: (value) {
                            setState(() {
                              weightControllerHasError = false;
                            });
                          },
                          decoration: getTextFieldDecoration(
                            weightControllerHasError,
                          ),
                          inputFormatters: [
                            FilteringTextInputFormatter.allow(RegExp(r'[0-9]'))
                          ],
                          controller: weightController,
                          keyboardType: TextInputType.number,
                          placeholder: "Weight (kg)",
                          style: TextStyle().copyWith(
                              color: Theme.of(context)
                                  .colorScheme
                                  .primaryContainer),
                        ),
                      ),
                    ],
                  ),
                ),
                SizedBox(width: 8),
                Expanded(
                  flex: 2,
                  child: CupertinoButton(
                    minSize: 0,
                    padding: const EdgeInsets.all(7.0),
                    borderRadius: BorderRadius.circular(7),
                    color: Theme.of(context).colorScheme.primary,
                    onPressed: () {
                      bool isValid = true;

                      final String repsString = repsController.text;
                      final String weightString = weightController.text;

                      if (repsString.isEmpty) {
                        repsControllerHasError = true;
                        isValid = false;
                      }
                      if (weightString.isEmpty) {
                        weightControllerHasError = true;
                        isValid = false;
                      }

                      if (!isValid) {
                        setState(() {});
                      } else {
                        final int reps = int.parse(repsController.text);
                        final int weight = int.parse(weightController.text);
                        final exerciseSet = ExerciseSetEntity(
                          weight: weight,
                          reps: reps,
                        );
                        widget.setAdderCallback(exerciseSet);
                      }

                      // Something like setState()
                    },
                    child: Text(
                      'Add Set',
                      style: TextStyle().copyWith(
                        color: Theme.of(context).colorScheme.secondaryContainer,
                      ),
                    ),
                  ),
                )
              ],
            )
          : SizedBox.shrink(),
    );
  }

  /// Get decoration for CupertinoStyled TextField
  BoxDecoration getTextFieldDecoration(bool hasError) {
    return BoxDecoration(
      color: CupertinoDynamicColor.withBrightness(
        color: CupertinoColors.white,
        darkColor: CupertinoColors.black,
      ),
      border: Border.all(
        color: CupertinoDynamicColor.withBrightness(
          color: hasError
              ? Theme.of(context).colorScheme.error
              : Color(0x33000000),
          darkColor: hasError
              ? Theme.of(context).colorScheme.error
              : Color(0x33FFFFFF),
        ),
        width: 0.0,
      ),
      borderRadius: BorderRadius.all(Radius.circular(5.0)),
    );
  }
}
